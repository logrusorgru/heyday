package heyday

import (
	"fmt" // DEBUG
	"math"
)

func (c *Lab) XYZ(io ...int) *XYZ {
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
	wt := WhitePoint(wp.X, wp.Y) // type XYZ
	wp = nil                     // GC (Is it really needed?)
	y := (c.L + 16) / 116        // f(Qy)
	x := c.A/500 + y             // f(Qx)
	z := y - c.B/200             // f(Qz)
	if x > c6d29 {               // eq.: Qx^3 > (6/29)^3 |=> Qx > 6/29
		x = math.Pow(x, 3) // Qx = f(Qx)^3
	} else {
		x = c108d841 * (x - c4d29)
	}
	if y > c6d29 {
		y = math.Pow(y, 3)
	} else {
		y = c108d841 * (y - c4d29)
	}
	if z > c6d29 {
		z = math.Pow(z, 3)
	} else {
		z = c108d841 * (z - c4d29)
	}
	x = x * wt.X // X = Qx*Xn (<=| Qx = (X/Xn))
	y = y * wt.Y // 100
	z = z * wt.Z
	xyz := &XYZ{x, y, z}
	return xyz
}

// Create CIE 1931 Color Space (CIE XYZ) from current CIE L*a*b*
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// ref.: ASTM E308 http://wenku.baidu.com/view/1dc90ac20c22590102029dce
//            math http://mathb.in/24271

func (c *Lab) Hue() float64 {
	return hue(c.A, c.B)
}

// return CIE 1976 hue angle [0, 360]° of current color point

func (c *Lab) Chromas() float64 {
	return chromas(c.A, c.B)
}

// return CIE 1976 chromas o current color point

func (c *Lab) LCH() *LCH {
	h := hue(c.A, c.B)
	cc := chromas(c.A, c.B)
	lch := &LCH{c.L, cc, h}
	return lch
}

// create L*CH° version of CIE L*,a*,b*

func (c *Lab) Lab() *Lab {
	return c
}

// return sefl

func (c *Lab) It() (l, a, b float64) {
	return c.L, c.A, c.B
}

// return self values

// DEBUG
func (c *Lab) Show() {
	fmt.Println("CIE L*,a*,b*")
	fmt.Printf("L*: %.48f\n", c.L)
	fmt.Printf("a*: %.48f\n", c.A)
	fmt.Printf("b*: %.48f\n", c.B)
}
