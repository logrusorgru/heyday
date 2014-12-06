package heyday

import (
	"fmt" // DEBUG
	. "github.com/logrusorgru/heyday"
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
	return c.Labl(&white_points[observer][illuminant])
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

func (c *XYZ) Labl(wp *WP) *Lab {
	wt := TristimulusWhite(wp.X, wp.Y) // type XYZ
	wp = nil                           // GC (Is it really needed?)
	qx := c.X / wt.X                   // Qx
	qy := c.Y / wt.Y                   // Qy
	qz := c.Z / wt.Z                   // Qz
	if qx > c6d29p3 {                  // (6/29)^3
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
	return &Lab{l, a, b}
}

// more flexible method
// you can use your own XYZ white point

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
	return c.Luvl(&white_points[observer][illuminant])
}

// Create CIE L*u*v* from current XYZ.
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// ref.: ASTM E308 http://wenku.baidu.com/view/1dc90ac20c22590102029dce

func (c *XYZ) Luvl(wp *WP) *Luv {
	wt := TristimulusWhite(wp.X, wp.Y) // type XYZ
	wp = nil                           // GC (Is it really needed?)
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
	return &Luv{l, u, v}
}

// more flexible method
// you can use your own XYZ white point

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
	return c.HunterLabl(&white_points[observer][illuminant])
}

// Create Hunter L,a,b from current XYZ.
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// ref.: http://fp.optics.arizona.edu/opti588/reading/CIE_Lab_color_space.pdf
//       http://en.wikipedia.org/wiki/Lab_color_space

func (c *XYZ) HunterLabl(wp *WP) *HunterLab {
	wt := TristimulusWhite(wp.X, wp.Y) // type XYZ
	wp = nil                           // GC (Is it really needed?)
	//
	qy := c.Y / wt.Y
	qx := c.X / wt.X
	qz := c.Z / wt.Z
	sqrt_qy := math.Sqrt(qy)
	l := 100 * sqrt_qy
	a := c175d198_04 * (wt.X + wt.Y) * (qx - qy) / sqrt_qy // Ka := c175d198_04 * (wt.X + wt.Y)
	b := c70d218_11 * (wt.Y + wt.Z) * (qy - qz) / sqrt_qy  // Kb := c70d218_11 * (wt.Y + wt.Z)
	//
	return &HunterLab{l, a, b}
}

// more flexible method
// you can use your own XYZ white point

func (c *XYZ) Yxy() *Yxy {
	x := c.X / (c.X + c.Y + c.Z)
	y := c.Y / (c.X + c.Y + c.Z)
	return &Yxy{x, y, c.Y}
}

// See also XYZ.Chromaticity() method
// ref.: https://en.wikipedia.org/wiki/CIE_1931_color_space#CIE_xy_chromaticity_diagram_and_the_CIE_xyY_color_space

func (c *XYZ) RGB(cio ...int) *RGB {
	var illuminant, observer, color_space int
	if len(cio) == 3 {
		observer = cio[2]
		illuminant = cio[1]
		color_space = cio[0]
	} else if len(cio) == 2 {
		illuminant = cio[1]
		color_space = cio[0]
		observer = O2
	} else if len(cio) == 1 {
		color_space = cio[0]
		illuminant = D65
		observer = O2
	} else {
		color_space = SRGB
		illuminant = D65
		observer = O2
	}
	rgb, _ := c.RGBl( // fuck the error
		&rgb_white_points[color_space],
		&white_points[observer][illuminant],
	)
	return rgb
}

// default RGB color space is sRGB
// default illuminant of XYZ white point is D65
// default observer of XYZ white point is CIE 1931 2° Observer
// call examples:
//  xyz.RGB()
//	xyz.RGB(Adobe_RGB_1998)
//	xyz.RGB(Adobe_RGB_1998, D50)
//	xyz.RGB(Adobe_RGB_1998, D50, O10)

func (c *XYZ) RGBl(color_space *Senary, white_point *WP) (*RGB, error) {
	imx, err := RgbInverseMatrix(
		color_space,
		white_point,
	)
	if err != nil {
		return nil, err
	}
	r, g, b := imx.RightColumn(c.X, c.Y, c.Z)
	return &RGB{r, g, b}, nil
}

// more flexible method
// you can use your own RGB color space and XYZ white point

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
