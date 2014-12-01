package heyday

import (
	"fmt" // DEBUG
	. "github.com/logrusorgru/heyday"
	"math"
)

func (c *RGB) XYZ(cio ...int) *XYZ {
	var illuminant, observer, color_space int
	if len(cio) == 3 {
		observer = cio[2]
		illuminant = cio[1]
		color_space = cio[0]
	} else if len(cio) == 2 {
		illuminant = cio[1]
		color_space = cio[0]
		observer = O2
	} else if len(cio) == 1 {
		color_space = cio[0]
		illuminant = D65
		observer = O2
	} else {
		color_space = SRGB
		illuminant = D65
		observer = O2
	}
	return c.XYZl(
		&rgb_white_points[color_space],
		&white_points[observer][illuminant],
	)
}

// default color space sRGB
// default illuminant D65
// default observer CIE 1931 2° Observer
// call examples:
//  rgb.XYZ()
//	rgb.XYZ(Adobe_RGB_1998)
//	rgb.XYZ(Adobe_RGB_1998, D50)
//	rgb.XYZ(Adobe_RGB_1998, D50, O10)

func (c *RGB) XYZl(color_space *Senary, white_point *WP) *XYZ {
	x, y, z := RgbDirectMatrix(
		color_space,
		white_point,
	).RightColumn(c.R, c.G, c.B)
	return &XYZ{x, y, z}
}

// more flexible method
// you can use your own RGB color space and XYZ white point

func (c *RGB) HSL() *HSL {
	// rgb [0,1]
	min := math.Min(c.R, math.Min(c.G, c.B))
	max := math.Max(c.R, math.Max(c.G, c.B))
	delta := max - min
	var h, s, l float64
	l = (max + min) / 2
	if delta == 0 {
		h = 0
		s = 0
	} else {
		if l < .5 {
			s = delta / (max + min)
		} else {
			s = delta / (2 - max - min)
		}
		deltaR := ((max-c.R)/6 + delta/2) / delta
		deltaG := ((max-c.G)/6 + delta/2) / delta
		deltaB := ((max-c.B)/6 + delta/2) / delta
		if c.R == max {
			h = deltaB - deltaG
		} else if c.G == max {
			h = 1./3 + deltaR - deltaB
		} else if c.B == max {
			h = 2./3 + deltaG - deltaR
		}
		if h < 0 {
			h += 1
		} else if h > 1 {
			h -= 1
		}
	}
	return &HSL{h, s, l}
}
