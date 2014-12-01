package heyday

import (
	"fmt" // DEBUG
	. "github.com/logrusorgru/heyday"
	"math"
)

func (c *HunterLab) XYZ(io ...int) *XYZ {
	var illuminant, observer int
	if len(io) == 2 {
		observer = io[1]
		illuminant = io[0]
	} else if len(io) == 1 {
		illuminant = io[0]
		observer = O2
	} else {
		illuminant = D65
		observer = O2
	}
	return c.XYZl(&white_points[observer][illuminant])
}

// Create CIE 1931 Color Space (CIE XYZ) from current Hunter L,a,b
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// ref.: http://fp.optics.arizona.edu/opti588/reading/CIE_Lab_color_space.pdf
//       http://en.wikipedia.org/wiki/Lab_color_space

func (c *HunterLab) XYZl(io ...int) *XYZ {
	wt := TristimulusWhite(wp.X, wp.Y) // type XYZ
	wp = nil                           // GC (Is it really needed?)
	qy := math.Pow(c.L/c100, 2)
	qx := c198_04d175*c.A*c.L*c1d100/(wt.X+wt.Y) + qy // c1d100 = 0.01 = 1/100 means /100
	qz := qy - c218_11d70*c.B*c.L*c1d100/(wt.Y+wt.Z)  //
	x := qx * wt.X
	y := qy * wt.Y
	z := qz * wt.Z
	return &XYZ{x, y, z}
}

// more flexible method
// you can use your own XYZ white point

func (c *HunterLab) HunterLab() *HunterLab {
	return c
}

// return sefl

// DEBUG
func (c *HunterLab) Show() {
	fmt.Println("Hunter L,a,b")
	fmt.Printf("L: %.48f\n", c.L)
	fmt.Printf("a: %.48f\n", c.A)
	fmt.Printf("b: %.48f\n", c.B)
}
