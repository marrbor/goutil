package geo

import "math"

type (
	JapaneseAddress struct {
		Pref  string `json:"pref"`
		City  string `json:"city"`
		Area  string `json:"area"`
		Block string `json:"block"`
	}
)

// String returns combined Japanese address string.
func (a *JapaneseAddress) String() string {
	return a.Pref + a.City + a.Area + a.Block
}

// IsValidLatitude returns whether given value is correct latitude value or not.
func IsValidLatitude(lat float64) bool {
	return lat >= -90 && lat <= 90
}

// IsValidLongitude returns whether given value is correct longitude value or not.
func IsValidLongitude(lon float64) bool {
	return lon >= -180 && lon <= 180
}

const (
	EquatorialRadius = 6378137.0            // radius of earth GRS80
	Eccentricity     = 0.081819191042815790 // GRS80
)

// HubenyDistance returns meter(s) of given two points. refer: http://qiita.com/tmnck/items/30b42ba5df28c38b0f89
func HubenyDistance(srcLatitude, srcLongitude, dstLatitude, dstLongitude float64) float64 {
	dx := (dstLongitude - srcLongitude) * math.Pi / 180
	dy := (dstLatitude - srcLatitude) * math.Pi / 180
	my := ((srcLatitude + dstLatitude) / 2) * math.Pi / 180

	W := math.Sqrt(1 - (math.Pow(Eccentricity, 2) * math.Pow(math.Sin(my), 2)))
	mNumer := EquatorialRadius * (1 - math.Pow(Eccentricity, 2))

	M := mNumer / math.Pow(W, 3)
	N := EquatorialRadius / W
	return math.Sqrt(math.Pow(dy*M, 2) + math.Pow(dx*N*math.Cos(my), 2))
}
