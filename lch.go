package heyday

import (
	"fmt" // DEBUG
	"math"
)

/*

	LCH Hue
	 [0, 360] or [-180, 180] ?
	 That is the fucking question.

*/

func (c *LCH) Lab() *Lab {
	a := math.Cos(degree_2_radian(c.H)) * math.Sqrt(c.C)
	b := math.Sin(degree_2_radian(c.H)) * math.Sqrt(c.C)
	lab := &Lab{c.L, a, b}
	return lab
}

// create CIE L*,a*,b* from current color point

func (c *LCH) Luv() *Luv {
	u := math.Cos(degree_2_radian(c.H)) * math.Sqrt(c.C)
	v := math.Sin(degree_2_radian(c.H)) * math.Sqrt(c.C)
	luv := &Luv{c.L, u, v}
	return luv
}

// create CIE L*,u*,v* from current color point

func (c *LCH) LCH() *LCH {
	return c
}

// return self

// DEBUG
func (c *LCH) Show() {
	fmt.Println("L*CH°")
	fmt.Printf("L*: %.48f", c.X)
	fmt.Printf("C : %.48f", c.Y)
	fmt.Printf("H°: %.48f", c.Z)
}
