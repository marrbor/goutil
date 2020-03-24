package enc_test

import (
	"testing"

	"github.com/marrbor/goutil/enc"
	"github.com/stretchr/testify/assert"
)

func TestEncryptPassword(t *testing.T) {
	t.Log(enc.Encrypt256Password("abcdefg"))
}
func TestHash32(t *testing.T) {
	n, err := enc.Hash32("abcdefg")
	assert.NoError(t, err)
	assert.EqualValues(t, uint32(0x2a9eb737), n)
}
