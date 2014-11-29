package heyday

import (
	"math"
)

func degree_2_radian(d float64) float64 {
	return (d / 180) * math.Pi
}

func radian_2_degree(r float64) float64 {
	return (r / math.Pi) * 180
}

func hue(a, b float64) float64 {
	h := math.Atan2(b, a)
	if h > 0 {
		return radian_2_degree(h)
	} else {
		//return 360 - radian_2_degree(math.Abs(h))
		return 360 + radian_2_degree(h)
	}
}

// hue from CIE L*,a*,b* or CIE L*,u*,v*

func chromas(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

// CIE 1976 chromas from CIE L*,a*,b* or CIE L*,u*,v*

func Round(f float64) int {
	return int(f + 0.5)
}

// round float64 to int (>=.5 ciel, <.5 floor)

func fround12(f float64) float64 {
	return float64(int64(f*1e12+0.5) / 1e12)
}

// round float64 to 12 digits after dot
// for lab deltas
