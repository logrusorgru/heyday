package heyday

import (
	"strings"
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
	rgb := &RGB{r, g, b}
	return rgb
}

// Windows
func (c *HSL) From240(h, s, l float64) *HSL {
	hsl := &HSL{
		h / 240,
		s / 240,
		l / 240,
	}
	return hsl
}

func (c *HSL) To240() (h, s, l float64) {
	h = c.H * 240
	s = c.S * 240
	l = c.L * 240
	return
}

// Paint Shop Pro
func (c *HSL) From255(h, s, l float64) *HSL {
	hsl := &HSL{
		h / 255,
		s / 255,
		l / 255,
	}
	return hsl
}

func (c *HSL) To255() (h, s, l float64) {
	h = c.H * 255
	s = c.S * 255
	l = c.L * 255
	return
}

// add alpha channel, create new HSLA

func (c *HSL) AddAlpha(a float64) *HSLA {
	hsla := &HSLA{
		c.H,
		c.S,
		c.L,
		a,
	}
	return hsla
}

// only HSL type have this method
// return string to css, like `hsl(128, 200, 50, .123)`
func (c *HSL) Css() string {
	fmt.Printf("hsl(%d, %d%%, %d%%)", round(c.H), round(c.S), round(c.L))
}

func (c *HSL) It() (h, l, s float64) {
	return c.H, c.L, c.S
}

func (c *HSL) HSL() *HSL {
	return c
}

// Deltas

func (c *HSL) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *HSL) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *HSL) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *HSL) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *HSL) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *HSL) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
