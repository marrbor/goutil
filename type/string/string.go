package string

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
)

// GetCode returns given length code generated from uuid.
func GetCode(number int) string {
	id, _ := uuid.NewRandom()
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

var randSrc = rand.NewSource(time.Now().UnixNano())

const (
	rsLetters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%^~*&+-=?_"
	rsPwLetters     = "abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789!#$%^~*&+-=?_"
	rsLetterIdxBits = 6
	rsLetterIdxMask = 1<<rsLetterIdxBits - 1
	rsLetterIdxMax  = 63 / rsLetterIdxBits
)

// RandString returns random string has given length.
func RandString(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), rsLetterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), rsLetterIdxMax
		}
		idx := int(cache & rsLetterIdxMask)
		if idx < len(rsLetters) {
			b[i] = rsLetters[idx]
			i--
		}
		cache >>= rsLetterIdxBits
		remain--
	}
	return string(b)
}

// PwString returns random string for password has given length.
func PwString(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), rsLetterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), rsLetterIdxMax
		}
		idx := int(cache & rsLetterIdxMask)
		if idx < len(rsPwLetters) {
			b[i] = rsPwLetters[idx]
			i--
		}
		cache >>= rsLetterIdxBits
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
