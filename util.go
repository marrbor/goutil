package goutil

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math"
	"math/rand"
	"net"
	"net/url"
	"path"
	"reflect"
	"runtime"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/sys/unix"
)

type (
	JapaneseAddress struct {
		Pref  string `json:"pref"`
		City  string `json:"city"`
		Area  string `json:"area"`
		Block string `json:"block"`
	}
)

// JST returns pointer of time.Location instance that points Japan Standard Time.
func JST() *time.Location {
	return time.FixedZone("Asia/Tokyo", 9*60*60)
}

// String returns combined Japanese address string.
func (a *JapaneseAddress) String() string {
	return a.Pref + a.City + a.Area + a.Block
}

// HasMapItem returns whether specified key is included specified map or not.
func HasMapItem(params map[string]interface{}, key string) bool {
	return params[key] != nil
}

// JSONString converts given structure to JSON string.
func JSONString(params interface{}) (string, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Hash32 converts string to hash value.
func Hash32(s string) (uint32, error) {
	h := fnv.New32a()
	if _, err := h.Write([]byte(s)); err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

// GetURLPathBase returns base(last part) of URL. ex '/api/v1/users/12345' => 12345
func GetURLPathBase(u *url.URL) string {
	return path.Base(u.Path)
}

// Encrypt256 returns SHA256 encrypted string.
func Encrypt256Password(src string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(src)))
}

// IsValidLatitude returns whether given value is correct latitude value or not.
func IsValidLatitude(lat float64) bool {
	return lat >= -90 && lat <= 90
}

// IsValidLongitude returns whether given value is correct longitude value or not.
func IsValidLongitude(lon float64) bool {
	return lon >= -180 && lon <= 180
}

// IsLastDayOfMonth returns whether given date is last day of month or not.
func IsLastDayOfMonth(time time.Time) bool {
	tomorrow := time.AddDate(0, 0, 1)
	return tomorrow.Day() == 1 // last day of month when next day is 1.
}

// IsFirstDayOfMonth returns whether given date is first day of month or not.
func IsFirstDayOfMonth(time time.Time) bool {
	return time.Day() == 1 // first day of month when 1.
}

// GetCode returns given length code generated from uuid.
func GetCode(number int) string {
	id, _ := uuid.NewV4()
	uuID := fmt.Sprintf("%s", id)
	if number <= 8 {
		return uuID[:number] // use first 8 digit.
	}

	u := strings.Replace(uuID, "-", "", -1)
	if number > len(u) {
		return "" // return empty when given digit over 32.
	}
	return u[:number]
}

// WaitSec waited given second(s).
func WaitSec(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}

// WaitMsec waited given millisecond(s).
func WaitMsec(msec int) {
	time.Sleep(time.Duration(msec) * time.Millisecond)
}

const (
	EquatorialRadius = 6378137.0            // radius of earth GRS80
	Eccentricity     = 0.081819191042815790 // GRS80
)

// HubenyDistance returns meter(s) of given two points. refer: http://qiita.com/tmnck/items/30b42ba5df28c38b0f89
func HubenyDistance(srcLatitude, srcLongitude, dstLatitude, dstLongitude float64) float64 {
	dx := (dstLongitude - srcLongitude) * math.Pi / 180
	dy := (dstLatitude - srcLatitude) * math.Pi / 180
	my := ((srcLatitude + dstLatitude) / 2) * math.Pi / 180

	W := math.Sqrt(1 - (math.Pow(Eccentricity, 2) * math.Pow(math.Sin(my), 2)))
	mNumer := EquatorialRadius * (1 - math.Pow(Eccentricity, 2))

	M := mNumer / math.Pow(W, 3)
	N := EquatorialRadius / W
	return math.Sqrt(math.Pow(dy*M, 2) + math.Pow(dx*N*math.Cos(my), 2))
}

var randSrc = rand.NewSource(time.Now().UnixNano())

const (
	rs6Letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rs6LetterIdxBits = 6
	rs6LetterIdxMask = 1<<rs6LetterIdxBits - 1
	rs6LetterIdxMax  = 63 / rs6LetterIdxBits
)

// RandString returns random string given length.
func RandString(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), rs6LetterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), rs6LetterIdxMax
		}
		idx := int(cache & rs6LetterIdxMask)
		if idx < len(rs6Letters) {
			b[i] = rs6Letters[idx]
			i--
		}
		cache >>= rs6LetterIdxBits
		remain--
	}
	return string(b)
}

// StructToStringMap convert structure contents to map[string]string.
// Field tag specified by tagName is used for key, and Field va string is
// Map keys are string that has been given value of tag specified by 'tagName'.
// Map values are value of that field, return empty map when nil is given for 's'.
func StructToStringMap(tagName string, s interface{}) *map[string]string {
	var ret = make(map[string]string)
	if s != nil {
		t := reflect.TypeOf(s)
		v := reflect.ValueOf(s)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get(tagName)
			ret[tag] = fmt.Sprintf("%v", v.FieldByName(field.Name))
		}
	}
	return &ret
}

// DetectOSVersion returns os version that run this program.
func DetectOSVersion() (string, error) {
	var ver string
	var err error
	os := runtime.GOOS
	switch os {
	case "windows":
		ver, err = getWindowsVer()
	case "darwin":
		fallthrough
	case "linux":
		ver, err = getUnixVer()
	default:
		err = fmt.Errorf("unsupport runtime")
	}
	return ver, err
}

func getWindowsVer() (string, error) {
	return "", nil
}

func uname2str(u []byte) string {
	str := ""
	for _, v := range u {
		m := int(v)
		if m <= 0 {
			break
		}
		str += string(m)
	}
	return str
}

func getUnixVer() (string, error) {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return "", err
	}
	return uname2str(uname.Version[:]), nil
}

func getUnixSysName() (string, error) {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return "", err
	}
	return uname2str(uname.Sysname[:]), nil
}

func getUnixNodeName() (string, error) {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return "", err
	}
	return uname2str(uname.Nodename[:]), nil
}

func getUnixMachineName() (string, error) {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return "", err
	}
	return uname2str(uname.Machine[:]), nil
}

func getUnixRelease() (string, error) {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		return "", err
	}
	return uname2str(uname.Release[:]), nil
}

// GetIP returns IP address of this reporter
func GetIP() (string, error) {
	// gather ip addresses.
	as, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	// seek ip address that is not a loopback.
	for _, a := range as {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil // got it.
			}
		}
	}
	return "", fmt.Errorf("effective adress not found")
}

// GetInterfaceName returns NIC for given IP.
func GetInterface(ip string) (*net.Interface, error) {
	// gather interfaces.
	ifs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// seek network interface bind for given IP
	for _, i := range ifs {
		if a, err := i.Addrs(); err == nil {
			for _, addr := range a {
				if strings.Contains(addr.String(), ip) {
					return &i, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("NIC not found")
}

// GetMacAddress returns mac address for given NIC.
func GetMacAddress(nic *net.Interface) (string, error) {
	// extract the hardware information base on the interface name capture above
	i, err := net.InterfaceByName(nic.Name)
	if err != nil {
		return "", err
	}
	hwa := i.HardwareAddr
	mac := hwa.String()
	return mac, nil
}
