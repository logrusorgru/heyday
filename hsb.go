package heyday

// just the same as hsv

import (
	"math"
)

func (c *HSB) RGB() *RGB {
	var r, g, b float64
	if c.S == 0 { //HSV from 0 to 1
		r = c.B * 255
		g = c.B * 255
		b = c.B * 255
	} else {
		h := c.H * 6
		if h == 6 {
			h = 0
		} //H must be < 1
		i := math.Floor(h) //Or ... var_i = floor( var_h )
		v1 := c.B * (1 - c.S)
		v2 := c.B * (1 - c.S*(h-i))
		v3 := c.B * (1 - c.S*(1-(h-i)))

		if i == 0 {
			r = c.B
			g = v3
			b = v1
		} else if i == 1 {
			r = v2
			g = c.B
			b = v1
		} else if i == 2 {
			r = v1
			g = c.B
			b = v3
		} else if i == 3 {
			r = v1
			g = v2
			b = c.B
		} else if i == 4 {
			r = v3
			g = v1
			b = c.B
		} else {
			r = c.B
			g = v1
			b = v2
		}

		r = r * 255 //RGB results from 0 to 255
		g = g * 255
		b = b * 255
	}
	rgb := &RGB{r, g, b}
	return rgb

}

// Gimp Photoshop Apple
func (c *HSB) From360_100(h, s, b float64) *HSB {
	hsv := &HSB{
		h / 360,
		s / 100,
		b / 100,
	}
	return hsv
}

func (c *HSB) To360_100() (h, s, b float64) {
	h = c.H * 360
	s = c.S * 100
	b = c.B * 100
	return
}

// Linux / KDE
func (c *HSB) From360_255(h, s, b float64) *HSB {
	hsv := &HSB{
		h / 360,
		s / 255,
		b / 255,
	}
	return hsv
}

func (c *HSB) To360_255() (h, s, b float64) {
	h = c.H * 360
	s = c.S * 255
	b = c.B * 255
	return
}

// GTK
func (c *HSB) From360_1(h, s, b float64) *HSB {
	hsv := &HSB{
		h / 360,
		s,
		b,
	}
	return hsv
}

func (c *HSB) To360_1() (h, s, b float64) {
	h = c.H * 360
	s = c.S
	b = c.B
	return
}

func (c *HSB) It() (h, s, b float64) {
	return c.H, c.S, c.B
}

func (c *HSB) HSB() *HSB {
	return c
}

// Deltas

func (c *HSB) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *HSB) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *HSB) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *HSB) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *HSB) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *HSB) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
