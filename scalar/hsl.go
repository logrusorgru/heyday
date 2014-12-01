package heyday

import (
	"fmt" // DEBUG
)

/*

http://www.easyrgb.com/index.php?X=MATH&H=22#text22

Applications		Space	H Range		S Range		L/V/B Range		Type

Paint Shop Pro		HSL		0 - 255		0 - 255		L	0 - 255		L1
Gimp				HSV		0 - 360°	0 - 100		V	0 - 100		V1
Photoshop			HSV		0 - 360°	0 - 100%	B	0 - 100%	V1
Windows				HSL		0 - 240		0 - 240		L	0 - 240		L2
Linux / KDE			HSV		0 - 360°	0 - 255		V	0 - 255		V2
GTK					HSV		0 - 360°	0 - 1.0		V	0 - 1.0		V3
Java (awt.Color)	HSV		0 - 1.0		0 - 1.0		B	0 - 1.0		Current
Apple				HSV		0 - 360°	0 - 100%	L	0 - 100%	V1

*/

func hue_2_rgb(v1, v2, vH float64) float64 { //Function Hue_2_RGB
	if vH < 0 {
		vH += 1
	}
	if vH > 1 {
		vH -= 1
	}
	if (6 * vH) < 1 {
		return (v1 + (v2-v1)*6*vH)
	}
	if (2 * vH) < 1 {
		return v2
	}
	if (3 * vH) < 2 {
		return (v1 + (v2-v1)*((2/3)-vH)*6)
	}
	return v1
}

// lost quality
// http://play.golang.org/p/9q5yBNDh3W

func (c *HSL) RGB() *RGB {
	var r, g, b float64
	if c.S == 0 { //HSL from 0 to 1
		r = c.L * 255 //RGB results from 0 to 255
		g = c.L * 255
		b = c.L * 255
	} else {
		var v1, v2 float64
		if c.L < 0.5 {
			v2 = c.L * (1 + c.S)
		} else {
			v2 = (c.L + c.S) - (c.S * c.L)
		}
		v1 = 2*c.L - v2
		r = 255 * hue_2_rgb(v1, v2, c.H+(1/3))
		g = 255 * hue_2_rgb(v1, v2, c.H)
		b = 255 * hue_2_rgb(v1, v2, c.H-(1/3))
	}
	return &RGB{r, g, b}
}
