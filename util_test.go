package goutil_test

import (
	"net"
	"net/url"
	"regexp"
	"testing"
	"time"

	"github.com/marrbor/goutil"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	input  interface{}
	expect interface{}
}

func TestJST(t *testing.T) {
	jst := goutil.JST()
	assert.EqualValues(t, "Asia/Tokyo", jst.String())
}

func TestJapaneseAddress_String(t *testing.T) {
	adr := goutil.JapaneseAddress{
		Pref:  "東京都",
		City:  "千代田区",
		Area:  "千代田",
		Block: "1-1-1",
	}
	assert.EqualValues(t, "東京都千代田区千代田1-1-1", adr.String())
}

func TestHasMapItem(t *testing.T) {
	tMap := map[string]interface{}{
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9,
	}
	for key := range tMap {
		assert.True(t, goutil.HasMapItem(tMap, key))
	}
	assert.False(t, goutil.HasMapItem(tMap, "ABC"))
	assert.False(t, goutil.HasMapItem(tMap, "DEF"))
	assert.False(t, goutil.HasMapItem(tMap, "012"))
	assert.False(t, goutil.HasMapItem(tMap, "!!!"))
}

func TestGetURLPathBase(t *testing.T) {
	var data = []testData{
		{input: "", expect: ""},
		{input: "/api/v1/xxx", expect: "xxx"},
		{input: "+81333573044", expect: ""},
		{input: "https://aaa/bbb/ccc", expect: "ccc"},
		{input: "https://aaa/bbb/ccc/", expect: "ccc"},
	}
	for _, entry := range data {
		urlStr, err := url.ParseRequestURI(entry.input.(string))
		if len(entry.expect.(string)) <= 0 {
			assert.Error(t, err)
			t.Logf("%s is not a valid url. skip testing", entry.input.(string))
			continue
		}
		base := goutil.GetURLPathBase(urlStr)
		assert.EqualValues(t, entry.expect, base)
	}
}

func TestIsValidLatitude(t *testing.T) {
	var data = []testData{
		{input: 90.001, expect: false},
		{input: -90.001, expect: false},
		{input: 90, expect: true},
		{input: -90, expect: true},
		{input: 89.999, expect: true},
		{input: -89.999, expect: true},
	}
	for _, entry := range data {
		lat, ok := entry.input.(float64)
		if !ok {
			lat = float64(lat)
		}
		assert.EqualValues(t, entry.expect, goutil.IsValidLatitude(lat))
	}
}

func TestIsValidLongitude(t *testing.T) {
	var data = []testData{
		{input: 180.001, expect: false},
		{input: -180.001, expect: false},
		{input: 180, expect: true},
		{input: -180, expect: true},
		{input: 179.999, expect: true},
		{input: -179.999, expect: true},
	}
	for _, entry := range data {
		lon, ok := entry.input.(float64)
		if !ok {
			lon = float64(lon)
		}
		assert.EqualValues(t, entry.expect, goutil.IsValidLongitude(lon))
	}
}

func TestIsLastDayOfMonth(t *testing.T) {
	var data = []testData{
		{input: time.Date(2016, 1, 30, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 1, 31, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 2, 1, 10, 00, 00, 0, time.Local), expect: false},

		// Leap Year
		{input: time.Date(2016, 2, 27, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 2, 28, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 2, 29, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 3, 1, 10, 00, 00, 0, time.Local), expect: false},

		// Normal Year
		{input: time.Date(2015, 2, 27, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2015, 2, 28, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2015, 3, 1, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 3, 30, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 3, 31, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 4, 1, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 4, 29, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 4, 30, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 5, 1, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 5, 30, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 5, 31, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 6, 1, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 6, 29, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 6, 30, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 7, 1, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 7, 30, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 7, 31, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 8, 1, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 8, 30, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 8, 31, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 9, 1, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 9, 29, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 9, 30, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 10, 1, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 10, 30, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 10, 31, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 11, 1, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 11, 29, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 11, 30, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 12, 1, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 12, 30, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 12, 31, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 1, 1, 10, 00, 00, 0, time.Local), expect: false},
	}
	for _, entry := range data {
		tm := entry.input.(time.Time)
		assert.EqualValues(t, entry.expect, goutil.IsLastDayOfMonth(tm))
	}
}

func TestIsFirstDayOfMonth(t *testing.T) {
	var data = []testData{
		{input: time.Date(2016, 1, 31, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 2, 1, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 2, 2, 10, 00, 00, 0, time.Local), expect: false},

		// leap year
		{input: time.Date(2016, 2, 29, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 3, 1, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 3, 2, 10, 00, 00, 0, time.Local), expect: false},

		// normal year
		{input: time.Date(2015, 2, 28, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2015, 3, 1, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2015, 3, 2, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 3, 31, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 4, 1, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 4, 2, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 4, 30, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 5, 1, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 5, 2, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 5, 31, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 6, 1, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 6, 2, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 6, 30, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 7, 1, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 7, 2, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 7, 31, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 8, 1, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 8, 2, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 8, 31, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 9, 1, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 9, 2, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 9, 30, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 10, 1, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 10, 2, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 10, 31, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 11, 1, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 11, 2, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 11, 30, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 12, 1, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 12, 2, 10, 00, 00, 0, time.Local), expect: false},

		{input: time.Date(2016, 12, 31, 10, 00, 00, 0, time.Local), expect: false},
		{input: time.Date(2016, 1, 1, 10, 00, 00, 0, time.Local), expect: true},
		{input: time.Date(2016, 1, 2, 10, 00, 00, 0, time.Local), expect: false},
	}
	for _, entry := range data {
		tm := entry.input.(time.Time)
		assert.EqualValues(t, entry.expect, goutil.IsFirstDayOfMonth(tm))
	}
}

func TestGetCode(t *testing.T) {
	for i := 0; i < 50; i++ {
		code := goutil.GetCode(i)
		t.Logf("Got code:%s\n", code)
		if i <= 32 {
			assert.EqualValues(t, i, len(code))
		} else {
			assert.EqualValues(t, 0, len(code))
		}
	}
}

func TestHubenyDistance(t *testing.T) {
	var data = []testData{
		{input: [][]float64{{35.123456, 135.123456}, {35.123456, 135.123456}}, expect: 0.0},
		{input: [][]float64{{-35.12345, 135.98765}, {-35.12345, 135.98765}}, expect: 0.0},
		{input: [][]float64{{35.1234, -135.9876}, {35.1234, -135.9876}}, expect: 0.0},
		{input: [][]float64{{-35.12, -135.98}, {-35.12, -135.98}}, expect: 0.0},
	}
	for _, entry := range data {
		lat1 := entry.input.([][]float64)[0][0]
		lon1 := entry.input.([][]float64)[0][1]
		lat2 := entry.input.([][]float64)[1][0]
		lon2 := entry.input.([][]float64)[1][1]
		assert.EqualValues(t, entry.expect, goutil.HubenyDistance(lat1, lon1, lat2, lon2))
	}
}

func TestEncryptPassword(t *testing.T) {
	t.Log(goutil.Encrypt256Password("abcdefg"))
}

func TestJSONString(t *testing.T) {
	type x struct {
		S string `json:"s"`
	}
	i := x{S: "hello"}
	ret, err := goutil.JSONString(i)
	assert.NoError(t, err)
	assert.EqualValues(t, "{\"s\":\"hello\"}", ret)
}

func TestHash32(t *testing.T) {
	n, err := goutil.Hash32("abcdefg")
	assert.NoError(t, err)
	assert.EqualValues(t, uint32(0x2a9eb737), n)
}

func TestWaitSec(t *testing.T) {
	t.Logf("wait 1 sec")
	goutil.WaitSec(1)
	t.Logf("done")
}

func TestWaitMsec(t *testing.T) {
	t.Logf("wait 10 msec")
	goutil.WaitMsec(10)
	t.Logf("done")
}

func TestRandString(t *testing.T) {
	s := goutil.RandString(0)
	assert.EqualValues(t, 0, len(s))

	r := regexp.MustCompile(`^[A-Za-z0-9!#$%^~*&+\-=?_]+$`)
	for i := 1; i <= 100; i++ {
		s = goutil.RandString(i)
		assert.EqualValues(t, i, len(s))
		assert.True(t, r.Match([]byte(s)))
	}
}

func TestPwString(t *testing.T) {
	s := goutil.PwString(0)
	assert.EqualValues(t, 0, len(s))

	r := regexp.MustCompile(`^[abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789!#$%^~*&+\-=?_]+$`)
	for i := 1; i <= 100; i++ {
		s = goutil.PwString(i)
		assert.EqualValues(t, i, len(s))
		assert.True(t, r.Match([]byte(s)))
	}
}

func TestStructToStringMap(t *testing.T) {
	t1 := struct {
		ID     string   `json:"id"`
		Number int      `json:"number"`
		Array  []string `json:"array"`
	}{
		ID:     "abc",
		Number: 123,
		Array:  []string{"ABC", "DEF", "GHI"},
	}

	m := goutil.StructToStringMap("json", t1)
	assert.NotNil(t, m)
	mm := *m
	assert.EqualValues(t, "abc", mm["id"])
	assert.EqualValues(t, "123", mm["number"])

	m = goutil.StructToStringMap("json", nil)
	assert.EqualValues(t, 0, len(*m))
}

func TestDetectOSVersion(t *testing.T) {
	ver, err := goutil.DetectOSVersion()
	assert.NoError(t, err)
	t.Log(ver)
}

func TestGetMacAddress(t *testing.T) {
	ip, err := goutil.GetIP()
	assert.NoError(t, err)

	nic, err := goutil.GetInterface(ip)
	assert.NoError(t, err)

	mac, err := goutil.GetMacAddress(nic)
	assert.NoError(t, err)
	t.Log(mac)
}

func TestValidateMacAddress(t *testing.T) {
	nics, err := net.Interfaces()
	assert.NoError(t, err)
	for _, nic := range nics {
		adr := nic.HardwareAddr.String()
		if 0 < len(adr) {
			t.Logf(adr)
			assert.True(t, goutil.ValidateMacAddress(adr))
		}
	}
}
