package heyday

import (
	"math"
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

// lost quality
// http://play.golang.org/p/9q5yBNDh3W

func (c *HSV) RGB() *RGB {
	var r, g, b float64
	if c.S == 0 { //HSV from 0 to 1
		r = c.V * 255
		g = c.V * 255
		b = c.V * 255
	} else {
		h := c.H * 6
		if h == 6 {
			h = 0
		} //H must be < 1
		i := math.Floor(h)
		v1 := c.V * (1 - c.S)
		v2 := c.V * (1 - c.S*(h-i))
		v3 := c.V * (1 - c.S*(1-(h-i)))
		if i == 0 {
			r = c.V
			g = v3
			b = v1
		} else if i == 1 {
			r = v2
			g = c.V
			b = v1
		} else if i == 2 {
			r = v1
			g = c.V
			b = v3
		} else if i == 3 {
			r = v1
			g = v2
			b = c.V
		} else if i == 4 {
			r = v3
			g = v1
			b = c.V
		} else {
			r = c.V
			g = v1
			b = v2
		}
		r = r * 255 //RGB results from 0 to 255
		g = g * 255
		b = b * 255
	}
	return &RGB{r, g, b}

}

// Gimp Photoshop Apple
func (c *HSV) From360_100(h, s, v float64) *HSV {
	hsv := &HSV{
		h / 360,
		s / 100,
		v / 100,
	}
	return hsv
}

func (c *HSV) To360_100() (h, s, v float64) {
	h = c.H * 360
	s = c.S * 100
	v = c.V * 100
	return
}

// Linux / KDE
func (c *HSV) From360_255(h, s, v float64) *HSV {
	hsv := &HSV{
		h / 360,
		s / 255,
		v / 255,
	}
	return hsv
}

func (c *HSV) To360_255() (h, s, v float64) {
	h = c.H * 360
	s = c.S * 255
	v = c.V * 255
	return
}

// GTK
func (c *HSV) From360_1(h, s, v float64) *HSV {
	hsv := &HSV{
		h / 360,
		s,
		v,
	}
	return hsv
}

func (c *HSV) To360_1() (h, s, v float64) {
	h = c.H * 360
	s = c.S
	v = c.V
	return
}

// return self values

func (c *HSV) It() (h, s, v float64) {
	return c.H, c.S, c.V
}

// return self

func (c *HSV) HSV() *HSV {
	return c
}

// Deltas

func (c *HSV) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *HSV) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *HSV) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *HSV) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *HSV) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *HSV) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
