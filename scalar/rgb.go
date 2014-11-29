package heyday

import (
	"fmt" // DEBUG
)

func (c *RGB) XYZ(io ...int) *XYZ {
	var illuminant, observer int
	if len(io) == 2 {
		observer = io[1]
		illuminant = io[0]
	} else if len(io) == 1 {
		illuminant = io[0]
		observer = O2
	} else {
		illuminant = D65
		observer = O2
	}
	wp := &white_points[observer][illuminant]
	wt := WhitePoint(wp.X, wp.Y) // type XYZ
	wp = nil                     // GC (Is it really needed?)
	// STUFF BELOW

	// STUFF ABOVE
	xyz := &XYZ{x, y, z}
	return xyz
}
