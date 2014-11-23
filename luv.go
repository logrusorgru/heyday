package heyday

import (
	"fmt" // DEBUG
	"math"
)

//
// CIE 1976 (L*, u*, v*) color space (??? Is it true? ???)
//

func (c *Luv) XYZ(io ...int) *XYZ {
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
	wp := &white_points[observer][illuminant]
	wt := WhitePoint(wp.X, wp.Y)                // type XYZ
	wp = nil                                    // GC (Is it really needed?)
	usn := 4 * wt.X / (wt.X + 15*wt.Y + 3*wt.Z) // uₙ'
	vsn := 9 * wt.Y / (wt.X + 15*wt.Y + 3*wt.Z) // vₙ'
	y := (c.L - 16) / 116                       // pre-f(Qy)
	if y > 6/29 {
		y = math.Pow(y, 3) // Qy
	} else {
		y = y*c108d841 - c4d29*c108d841 // Qy
	}
	y = y * wt.Y // Qy = Y/Yₙ |=> Y = Qy * Yₙ
	x := (usn / vsn) * c9d4 * y
	z := (3/vsn - c3d4*(usn/vsn) - 5) * y
	xyz := &XYZ{x, y, z}
	return xyz
}

// Create CIE 1931 Color Space (CIE XYZ) from current CIE L*,u*,v*
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// ref.: ASTM E308 http://wenku.baidu.com/view/1dc90ac20c22590102029dce

func (c *Luv) Hue() float64 {
	return hue(c.A, c.B)
}

// return CIE 1976 hue angle [0, 360]° of current color point

func (c *Lab) Chromas() float64 {
	return chromas(c.A, c.B)
}

// return CIE 1976 chromas o current color point

func (c *Luv) LCH() *LCH {
	h := hue(c.U, c.V)
	c := chromas(c.U, c.V)
	lch := &LCH{c.L, c, h}
	return lch
}

// create L*CH° version of CIE L*,a*,b*

func (c *Luv) Luv() *Luv {
	return c
}

// return self

// DEBUG
func (c *Luv) Show() {
	fmt.Println("CIE L*,u*,v*")
	fmt.Printf("L*: %.48f", c.X)
	fmt.Printf("u*: %.48f", c.Y)
	fmt.Printf("v*: %.48f", c.Z)
}
