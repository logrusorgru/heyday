package heyday

import (
	"fmt" // DEBUG
	. "github.com/logrusorgru/heyday"
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
	return c.XYZl(&white_points[observer][illuminant])
}

// Create CIE 1931 Color Space (CIE XYZ) from current CIE L*,u*,v*
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// ref.: ASTM E308 http://wenku.baidu.com/view/1dc90ac20c22590102029dce

func (c *Luv) XYZl(wp *WP) *XYZ {
	wt := TristimulusWhite(wp.X, wp.Y)        // type XYZ
	wp = nil                                  // GC (Is it really needed?)
	usn := 4 * wt.X / (wt.X + c1500 + 3*wt.Z) // uₙ'
	vsn := 9 * wt.Y / (wt.X + c1500 + 3*wt.Z) // vₙ'
	y := (c.L + 16) / 116                     // pre-f(Qy)
	if y > c6d29cbrt {
		y = math.Pow(y, 3) // Qy
	} else {
		y = c108d841 * (y - c4d29) // Qy
	}
	us := c.U/(13*c.L) + usn
	vs := c.V/(13*c.L) + vsn
	y = y * wt.Y // Qy = Y/Yₙ |=> Y = Qy * Yₙ
	x := -(9 * y * us) / ((us-4)*vs - us*vs)
	z := (9*y - 15*vs*y - vs*x) / (3 * vs)
	return &XYZ{x, y, z}
}

// more flexible method
// you can use your own XYZ white point

func (c *Luv) Hue() float64 {
	return hue(c.U, c.V)
}

// return CIE 1976 hue angle [0, 360]° of current color point

func (c *Luv) Chromas() float64 {
	return chromas(c.U, c.V)
}

// return CIE 1976 chromas o current color point

func (c *Luv) LCHuv() *LCHuv {
	h := hue(c.U, c.V)
	cc := chromas(c.U, c.V)
	return &LCHuv{c.L, cc, h}
}

// create L*CH°(uv) version of CIE L*,a*,b*

func (c *Luv) Luv() *Luv {
	return c
}

// return self

// DEBUG
func (c *Luv) Show() {
	fmt.Println("CIE L*,u*,v*")
	fmt.Printf("L*: %.48f\n", c.L)
	fmt.Printf("u*: %.48f\n", c.U)
	fmt.Printf("v*: %.48f\n", c.V)
}
