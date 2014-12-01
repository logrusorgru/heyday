package heyday

import (
	"fmt" // DEBUG
	. "github.com/logrusorgru/heyday"
	"math"
)

func (c HunterLabs) XYZs(io ...int) XYZs {
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
	return c.XYZsl(&white_points[observer][illuminant])
}

// Transform to CIE 1931 Color Space (CIE XYZ) from current Hunter L,a,b
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// ref.: http://fp.optics.arizona.edu/opti588/reading/CIE_Lab_color_space.pdf
//       http://en.wikipedia.org/wiki/Lab_color_space

func (c HunterLabs) XYZsl(wp *WP) XYZs {
	wt := TristimulusWhite(wp.X, wp.Y) // type XYZ
	wp = nil                           // GC (Is it really needed?)
	// main loop
	for i := 0; i < len(c); i++ {
		qy := math.Pow(c.L/c100, 2)
		qx := c198_04d175*c.A*c.L*c1d100/(wt.X+wt.Y) + qy // c1d100 = 0.01 = 1/100 means /100
		qz := qy - c218_11d70*c.B*c.L*c1d100/(wt.Y+wt.Z)  //
		c[i].L = qx * wt.X                                // X
		c[i].A = qy * wt.Y                                // Y
		c[i].B = qz * wt.Z                                // Z
	}
	// TRANSFORM Labs to XYZs
	cc := *(*XYZs)(unsafe.Pointer(&c))
	return cc
}

// more flexible method
// you can use your own XYZ white point

func (c HunterLabs) HunterLabs() HunterLabs {
	return c
}

// return sefl

// DEBUG
func (c HunterLabs) Show() {
	fmt.Println("Hunter L,a,b Array, Length %d", len(c))
	for i := 0; i < len(c); i++ {
		fmt.Printf("L : %.48f, a : %.48f, b : %.48f\n", c[i].L, c[i].A, c[i].B)
	}
}
