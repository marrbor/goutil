package string_test

import (
	"regexp"
	"testing"

	ms "github.com/marrbor/goutil/type/string"
	"github.com/stretchr/testify/assert"
)

func TestGetCode(t *testing.T) {
	for i := 0; i < 50; i++ {
		code := ms.GetCode(i)
		t.Logf("Got code:%s\n", code)
		if i <= 32 {
			assert.EqualValues(t, i, len(code))
		} else {
			assert.EqualValues(t, 0, len(code))
		}
	}
}

func TestRandString(t *testing.T) {
	s := ms.RandString(0)
	assert.EqualValues(t, 0, len(s))

	r := regexp.MustCompile(`^[A-Za-z0-9!#$%^~*&+\-=?_]+$`)
	for i := 1; i <= 100; i++ {
		s = ms.RandString(i)
		assert.EqualValues(t, i, len(s))
		assert.True(t, r.Match([]byte(s)))
	}
}

func TestPwString(t *testing.T) {
	s := ms.PwString(0)
	assert.EqualValues(t, 0, len(s))

	r := regexp.MustCompile(`^[abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789!#$%^~*&+\-=?_]+$`)
	for i := 1; i <= 100; i++ {
		s = ms.PwString(i)
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

	m := ms.StructToStringMap("json", t1)
	assert.NotNil(t, m)
	mm := *m
	assert.EqualValues(t, "abc", mm["id"])
	assert.EqualValues(t, "123", mm["number"])

	m = ms.StructToStringMap("json", nil)
	assert.EqualValues(t, 0, len(*m))
}
