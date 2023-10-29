package http_test

import (
	"net/url"
	"testing"

	"github.com/marrbor/goutil/net/http"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	input  interface{}
	expect interface{}
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
		base := http.GetURLPathBase(urlStr)
		assert.EqualValues(t, entry.expect, base)
	}
}
