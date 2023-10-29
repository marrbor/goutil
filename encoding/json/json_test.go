package json_test

import (
	"testing"

	"github.com/marrbor/goutil/encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestJSONString(t *testing.T) {
	type x struct {
		S string `json:"s"`
	}
	i := x{S: "hello"}
	ret, err := json.JSONString(i)
	assert.NoError(t, err)
	assert.EqualValues(t, "{\"s\":\"hello\"}", ret)
}
