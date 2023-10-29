package http_test

import (
	"testing"

	"github.com/marrbor/goutil/net/http"
	"github.com/stretchr/testify/assert"
)

func TestHTTPMethod(t *testing.T) {
	var x http.Method
	x = http.Methods.GET
	assert.EqualValues(t, "GET", x.String())
	x = http.Methods.HEAD
	assert.EqualValues(t, "HEAD", x.String())
	x = http.Methods.POST
	assert.EqualValues(t, "POST", x.String())
	x = http.Methods.PUT
	assert.EqualValues(t, "PUT", x.String())
	x = http.Methods.PATCH
	assert.EqualValues(t, "PATCH", x.String())
	x = http.Methods.DELETE
	assert.EqualValues(t, "DELETE", x.String())
	x = http.Methods.CONNECT
	assert.EqualValues(t, "CONNECT", x.String())
	x = http.Methods.OPTIONS
	assert.EqualValues(t, "OPTIONS", x.String())
	x = http.Methods.TRACE
	assert.EqualValues(t, "TRACE", x.String())
}
