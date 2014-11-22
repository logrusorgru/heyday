package heyday

import "math"

func deg2rad(d float64) float64 {
	return (d / 180) * math.Pi
}

func rad2deg(r float64) float64 {
	return (r / math.Pi) * 180
}

func round(f float64) int {
	return int(f + 0.5)
}
