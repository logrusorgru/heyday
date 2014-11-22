package heyday

import (
	"fmt"
)

// remove alpha channel from color, create new RGB color

func (c *RGBA) RmAlpha() *RGB {
	rgb := &RGB{
		c.R,
		c.G,
		c.B,
	}
	return rgb
}

// only this type (RGBA) have this method
// return string to css, like `rgba(128, 200, 50, .123)`
func (c *RGBA) Css() string {
	alpha := fmt.Sprintf("%g", float64(round(c.A*1000))/1000) // 1000 is precision
	fmt.Printf("rgba(%d, %d, %d, %s)", round(c.R), round(c.G), round(c.B), alpha[1:])
}

// return self vales

func (c *RGBA) It() (r, g, b, a float64) {
	return c.R, c.G, c.B, c.A
}

// return self

func (c *RGBA) RGBA() *RGBA {
	return c
}

// Deltas

func (c *RGBA) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *RGBA) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *RGBA) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *RGBA) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *RGBA) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *RGBA) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
