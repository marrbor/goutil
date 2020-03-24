package time_test

import (
	"testing"

	"github.com/marrbor/goutil/time"
	"github.com/stretchr/testify/assert"
)

func TestJST(t *testing.T) {
	jst := time.JST()
	assert.EqualValues(t, "Asia/Tokyo", jst.String())
}

func TestWaitSec(t *testing.T) {
	t.Logf("wait 1 sec")
	time.WaitSec(1)
	t.Logf("done")
}

func TestWaitMsec(t *testing.T) {
	t.Logf("wait 10 msec")
	time.WaitMsec(10)
	t.Logf("done")
}
