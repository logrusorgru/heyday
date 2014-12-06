package heyday

import (
	"fmt" // DEBUG
	. "github.com/logrusorgru/heyday"
	"math"
)

/*

	LCH Hue
	 [0, 360] or [-180, 180] ?
	 That is the fucking question.

*/

func (c *LCHuv) Luv() *Luv {
	u := math.Cos(degree_2_radian(c.H)) * c.C
	v := math.Sin(degree_2_radian(c.H)) * c.C
	return &Luv{c.L, u, v}
}

// create CIE L*,u*,v* from current color point

func (c *LCHuv) LCHuv() *LCHuv {
	return c
}

// return self

// DEBUG
func (c *LCHuv) Show() {
	fmt.Println("L*CH°uv")
	fmt.Printf("L*: %.48f\n", c.L)
	fmt.Printf("C : %.48f\n", c.C)
	fmt.Printf("H°: %.48f\n", c.H)
}
