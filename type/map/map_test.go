package maputil_test

import (
	"testing"

	maputil "github.com/marrbor/goutil/type/map"
	"github.com/stretchr/testify/assert"
)

func TestHasMapItem(t *testing.T) {
	tMap := map[string]interface{}{
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9,
	}
	for key := range tMap {
		assert.True(t, maputil.HasMapItem(tMap, key))
	}
	assert.False(t, maputil.HasMapItem(tMap, "ABC"))
	assert.False(t, maputil.HasMapItem(tMap, "DEF"))
	assert.False(t, maputil.HasMapItem(tMap, "012"))
	assert.False(t, maputil.HasMapItem(tMap, "!!!"))
}
