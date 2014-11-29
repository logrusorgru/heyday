package heyday

import (
	"fmt" // DEBUG
	"math"
)

func (c *XYZ) Chromaticity() (x, y float64) {
	x = c.X / (c.X + c.Y + c.Z)
	y = c.Y / (c.X + c.Y + c.Z)
	// z = 1 - x - y
	return
}

// get chromaticity co-ordinates x, y of current color

func (c *XYZ) Lab(io ...int) *Lab {
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
	qx := c.X / wt.X             // Qx
	qy := c.Y / wt.Y             // Qy
	qz := c.Z / wt.Z             // Qz
	if qx > c6d29p3 {            // (6/29)^3
		qx = math.Cbrt(qx) // f(Qx) = Qx^(1/3) |=> Qx = Cbrt( f(Qx) )
	} else {
		qx = c841d108*qx + c4d29 // f(Qx) = (841/108)*Qx + 4/29
	}
	if qy > c6d29p3 {
		qy = math.Cbrt(qy)
	} else {
		qy = c841d108*qy + c4d29
	}
	if qz > c6d29p3 {
		qz = math.Cbrt(qz)
	} else {
		qz = c841d108*qz + c4d29
	}
	l := 116*qy - 16     // L* = 116*f(Qy) -16
	a := 500 * (qx - qy) // a* = 500( f(Qx) - f(Qy) )
	b := 200 * (qy - qz) // b* = 200*( f(Qy) - f(Qz) )
	lab := &Lab{l, a, b}
	return lab
}

// Create CIE L*a*b* from current XYZ.
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// For examples:
// lab := xyz.Lab(D50, O10)
// lab := xyz.Lab()			// defaults
// lab := xyz.Lab(F11)		// CIE 1931 2° Observer (default)
// ref.: ASTM E308 http://wenku.baidu.com/view/1dc90ac20c22590102029dce

func (c *XYZ) Luv(io ...int) *Luv {
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
	qy := c.Y / wt.Y
	if qy > c6d29p3 {
		qy = math.Cbrt(qy)
	} else {
		qy = c841d108*qy + c4d29
	}
	us := 4 * c.X / (c.X + 15*c.Y + 3*c.Z)
	vs := 9 * c.Y / (c.X + 15*c.Y + 3*c.Z)
	usn := 4 * wt.X / (wt.X + c1500 + 3*wt.Z) // 1500 = 15 * wt.Y
	vsn := 9 * wt.Y / (wt.X + c1500 + 3*wt.Z) //
	l := 116*qy - 16
	u := 13 * l * (us - usn)
	v := 13 * l * (vs - vsn)
	luv := &Luv{l, u, v}
	return luv
}

// Create CIE L*u*v* from current XYZ.
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// ref.: ASTM E308 http://wenku.baidu.com/view/1dc90ac20c22590102029dce

func (c *XYZ) HunterLab(io ...int) *HunterLab {
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
	//
	qy := c.Y / wt.Y
	qx := c.X / wt.X
	qz := c.Z / wt.Z
	sqrt_qy := math.Sqrt(qy)
	l := 100 * sqrt_qy
	a := c175d198_04 * (wt.X + wt.Y) * (qx - qy) / sqrt_qy // Ka := c175d198_04 * (wt.X + wt.Y)
	b := c70d218_11 * (wt.Y + wt.Z) * (qy - qz) / sqrt_qy  // Kb := c70d218_11 * (wt.Y + wt.Z)
	//
	hunter_lab := &HunterLab{l, a, b}
	return hunter_lab
}

// Create Hunter L,a,b from current XYZ.
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// ref.: http://fp.optics.arizona.edu/opti588/reading/CIE_Lab_color_space.pdf
//       http://en.wikipedia.org/wiki/Lab_color_space

func (c *XYZ) Yxy() *Yxy {
	x := c.X / (c.X + c.Y + c.Z)
	y := c.Y / (c.X + c.Y + c.Z)
	yxy := &Yxy{x, y, c.Y}
	return yxy
}

// See also XYZ.Chromaticity() method
// ref.: https://en.wikipedia.org/wiki/CIE_1931_color_space#CIE_xy_chromaticity_diagram_and_the_CIE_xyY_color_space

func (c *XYZ) XYZ() *XYZ {
	return c
}

// return self

// DEBUG
func (c *XYZ) Show() {
	fmt.Println("CIE XYZ")
	fmt.Printf("X: %.48f\n", c.X)
	fmt.Printf("Y: %.48f\n", c.Y)
	fmt.Printf("Z: %.48f\n", c.Z)
}
