package time_test

import (
	"testing"
	"time"

	mt "github.com/marrbor/goutil/time"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	input  interface{}
	expect interface{}
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
		assert.EqualValues(t, entry.expect, mt.IsLastDayOfMonth(tm))
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
		assert.EqualValues(t, entry.expect, mt.IsFirstDayOfMonth(tm))
	}
}
