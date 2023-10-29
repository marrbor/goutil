package geo_test

import (
	"testing"

	"github.com/marrbor/goutil/geo"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	input  interface{}
	expect interface{}
}

func TestJapaneseAddress_String(t *testing.T) {
	adr := geo.JapaneseAddress{
		Pref:  "東京都",
		City:  "千代田区",
		Area:  "千代田",
		Block: "1-1-1",
	}
	assert.EqualValues(t, "東京都千代田区千代田1-1-1", adr.String())
}

func TestIsValidLatitude(t *testing.T) {
	var data = []testData{
		{input: 90.001, expect: false},
		{input: -90.001, expect: false},
		{input: 90, expect: true},
		{input: -90, expect: true},
		{input: 89.999, expect: true},
		{input: -89.999, expect: true},
	}
	for _, entry := range data {
		lat, _ := entry.input.(float64)
		assert.EqualValues(t, entry.expect, geo.IsValidLatitude(lat))
	}
}

func TestIsValidLongitude(t *testing.T) {
	var data = []testData{
		{input: 180.001, expect: false},
		{input: -180.001, expect: false},
		{input: 180, expect: true},
		{input: -180, expect: true},
		{input: 179.999, expect: true},
		{input: -179.999, expect: true},
	}
	for _, entry := range data {
		lon, _ := entry.input.(float64)
		assert.EqualValues(t, entry.expect, geo.IsValidLongitude(lon))
	}
}

func TestHubenyDistance(t *testing.T) {
	var data = []testData{
		{input: [][]float64{{35.123456, 135.123456}, {35.123456, 135.123456}}, expect: 0.0},
		{input: [][]float64{{-35.12345, 135.98765}, {-35.12345, 135.98765}}, expect: 0.0},
		{input: [][]float64{{35.1234, -135.9876}, {35.1234, -135.9876}}, expect: 0.0},
		{input: [][]float64{{-35.12, -135.98}, {-35.12, -135.98}}, expect: 0.0},
	}
	for _, entry := range data {
		lat1 := entry.input.([][]float64)[0][0]
		lon1 := entry.input.([][]float64)[0][1]
		lat2 := entry.input.([][]float64)[1][0]
		lon2 := entry.input.([][]float64)[1][1]
		assert.EqualValues(t, entry.expect, geo.HubenyDistance(lat1, lon1, lat2, lon2))
	}
}
