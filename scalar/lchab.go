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

func (c *LCHab) Lab() *Lab {
	a := math.Cos(degree_2_radian(c.H)) * c.C
	b := math.Sin(degree_2_radian(c.H)) * c.C
	lab := &Lab{c.L, a, b}
	return lab
}

// create CIE L*,a*,b* from current color point

func (c *LCHab) LCHab() *LCHab {
	return c
}

// return self

// DEBUG
func (c *LCH) Show() {
	fmt.Println("L*CH°ab")
	fmt.Printf("L*: %.48f\n", c.L)
	fmt.Printf("C : %.48f\n", c.C)
	fmt.Printf("H°: %.48f\n", c.H)
}
