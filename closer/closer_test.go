package closer_test

import (
	"os"
	"testing"

	"github.com/marrbor/goutil/closer"
	"github.com/stretchr/testify/assert"
)

func TestClose(t *testing.T) {
	f, err := os.Open("./closer_test.go")
	assert.NoError(t, err)
	closer.Close(f)
}
