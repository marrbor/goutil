package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUname(t *testing.T) {
	res, err := getUnixSysName()
	assert.NoError(t, err)
	t.Log(res)

	res, err = getUnixNodeName()
	assert.NoError(t, err)
	t.Log(res)

	res, err = getUnixMachineName()
	assert.NoError(t, err)
	t.Log(res)

	res, err = getUnixRelease()
	assert.NoError(t, err)
	t.Log(res)
}
