package heyday

import (
	"math"
)

func (c *LCH) LAB() *LAB {
	l := c.L
	a := math.Cos(degree_2_radian(c.H)) * c.C
	b := math.Sin(degree_2_radian(c.H)) * c.C
	lab := &LAB{l, a, b}
	return lab
}

func (c *LCH) LUV() *LUV {
	l := c.L
	u := c.C * math.Cos(c.H)
	v := c.C * math.Sin(c.H)
	luv := &LUV{l, u, v}
	return luv
}

// return self values

func (c *LCH) It() (l, cc, h float64) {
	return c.L, c.C, c.H
}

// return self

func (c *LCH) LCH() *LCH {
	return c
}

// Deltas

func (c *LCH) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *LCH) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *LCH) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *LCH) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *LCH) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *LCH) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
