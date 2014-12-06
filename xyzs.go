package heyday

import (
	"fmt" // DEBUG
	. "github.com/logrusorgru/heyday"
	"math"
	"unsafe"
)

func (c XYZs) Chromaticity() []CC {
	t := []CC{}
	for i := 0; i < len(c); i++ {
		x := c.X / (c.X + c.Y + c.Z)
		y := c.Y / (c.X + c.Y + c.Z)
		t = append(t, CC{x, y})
	}
	return t
}

// get chromaticity co-ordinates x, y of current colors

func (c XYZs) Labs(io ...int) Labs {
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
	return c.Labsl(&white_points[observer][illuminant])
}

// Transform to CIE L*a*b*'s from current XYZ's.
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// For examples:
// labs := xyzs.Lab(D50, O10)
// labs := xyzs.Lab()			// defaults
// labs := xyzs.Lab(F11)		// CIE 1931 2° Observer (default)
// ref.: ASTM E308 http://wenku.baidu.com/view/1dc90ac20c22590102029dce

func (c XYZs) Labsl(wp *WP) Labs {
	wt := TristimulusWhite(wp.X, wp.Y) // type XYZ
	wp = nil                           // GC (Is it really needed?)
	// main loop
	for i := 0; i < len(c); i++ {
		qx := c[i].X / wt.X // Qx
		qy := c[i].Y / wt.Y // Qy
		qz := c[i].Z / wt.Z // Qz
		if qx > c6d29p3 {   // (6/29)^3
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
		/*
			l := 116*qy - 16     // L* = 116*f(Qy) -16
			a := 500 * (qx - qy) // a* = 500( f(Qx) - f(Qy) )
			b := 200 * (qy - qz) // b* = 200*( f(Qy) - f(Qz) )
			c[i].X = l
			c[i].Y = a
			c[i].Z = b
		*/
		c[i].X = 116*qy - 16     // L* = 116*f(Qy) -16
		c[i].Y = 500 * (qx - qy) // a* = 500( f(Qx) - f(Qy) )
		c[i].Z = 200 * (qy - qz) // b* = 200*( f(Qy) - f(Qz) )
	}
	// TRANSFORM XYZs to Labs
	cc := *(*Labs)(unsafe.Pointer(&c))
	// Now c is the XYZs, but cc is the Labs with the same values
	// It is just the same pointer with another type
	// Fuck yeahhh!!!
	return cc
}

// more flexible method
// you can use your own XYZ white point

func (c XYZs) Luvs(io ...int) Luvs {
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
	return c.Luvsl(&white_points[observer][illuminant])
}

// Transform to CIE L*u*v*'s' from current XYZ's.
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// ref.: ASTM E308 http://wenku.baidu.com/view/1dc90ac20c22590102029dce

func (c XYZs) Luvsl(wp *WP) Luvs {
	wt := TristimulusWhite(wp.X, wp.Y) // type XYZ
	wp = nil                           // GC (Is it really needed?)
	// main loop
	for i := 0; i < len(c); i++ {
		qy := c[i].Y / wt.Y
		if qy > c6d29p3 {
			qy = math.Cbrt(qy)
		} else {
			qy = c841d108*qy + c4d29
		}
		us := 4 * c[i].X / (c[i].X + 15*c[i].Y + 3*c[i].Z)
		vs := 9 * c[i].Y / (c[i].X + 15*c[i].Y + 3*c[i].Z)
		usn := 4 * wt.X / (wt.X + c1500 + 3*wt.Z) // 1500 = 15 * wt.Y
		vsn := 9 * wt.Y / (wt.X + c1500 + 3*wt.Z) //
		c[i].X = 116*qy - 16
		c[i].Y = 13 * l * (us - usn)
		c[i].Z = 13 * l * (vs - vsn)
	}
	// TRANSFORM XYZs to Luvs
	cc := *(*Luvs)(unsafe.Pointer(&c))
	return cc
}

// more flexible method
// you can use your own XYZ white point

func (c XYZs) HunterLabs(io ...int) HunterLabs {
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
	return c.HunterLabsl(&white_points[observer][illuminant])
}

// Transform to Hunter L,a,b's from current XYZ's.
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// ref.: http://fp.optics.arizona.edu/opti588/reading/CIE_Lab_color_space.pdf
//       http://en.wikipedia.org/wiki/Lab_color_space

func (c XYZs) HunterLabsl(wp *WP) HunterLabs {
	wt := TristimulusWhite(wp.X, wp.Y) // type XYZ
	wp = nil                           // GC (Is it really needed?)
	// main loop
	for i := 0; i < len(c); i++ {
		qy := c[i].Y / wt.Y
		qx := c[i].X / wt.X
		qz := c[i].Z / wt.Z
		sqrt_qy := math.Sqrt(qy)
		c[i].Y = 100 * sqrt_qy
		c[i].X = c175d198_04 * (wt.X + wt.Y) * (qx - qy) / sqrt_qy // Ka := c175d198_04 * (wt.X + wt.Y)
		c[i].Z = c70d218_11 * (wt.Y + wt.Z) * (qy - qz) / sqrt_qy  // Kb := c70d218_11 * (wt.Y + wt.Z)
	}
	//
	// TRANSFORM XYZs to HunterLabs
	cc := *(*HunterLabs)(unsafe.Pointer(&c))
	return cc
}

// more flexible method
// you can use your own XYZ white point

func (c XYZs) Yxys() Yxys {
	for i := 0; i < len(c); i++ {
		x := c[i].X / (c[i].X + c[i].Y + c[i].Z)
		y := c[i].Y / (c[i].X + c[i].Y + c[i].Z) // Y2
		c[i].X = x
		c[i].Z = y // Y2
	}
	// TRANSFORM XYZs to Yxys
	cc := *(*Yxys)(unsafe.Pointer(&c))
	// Now c is the XYZs, but cc is the Yxys with the same values
	// It is just the same pointer with another type
	return cc
}

// See also XYZs.Chromaticity() method
// ref.: https://en.wikipedia.org/wiki/CIE_1931_color_space#CIE_xy_chromaticity_diagram_and_the_CIE_xyY_color_space

func (c XYZs) RGBs(cio ...int) RGBs {
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
	rgbs, _ := c.RGBsl(
		&rgb_white_points[color_space],
		&white_points[observer][illuminant],
	)
	return rgbs
}

// default color space sRGB
// default illuminant D65
// default observer CIE 1931 2° Observer
// call examples:
//  xyzs.RGBs()                          | sRGB, D65, CIE 1931 2° Observer
//	xyzs.RGBs(Adobe_RGB_1998)            | Adobe_RGB_1998, D65, CIE 1931 2° Observer
//	xyzs.RGBs(Adobe_RGB_1998, D50)       | Adobe_RGB_1998, D50, CIE 1931 2° Observer
//	xyzs.RGBs(Adobe_RGB_1998, D50, O10)  | Adobe_RGB_1998, D50, CIE 1964 10° Observer

func (c XYZs) RGBsl(color_space *Senary, white_point *WP) (RGBs, error) {
	imx, err := RgbInverseMatrix(
		color_space,
		white_point,
	)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(c); i++ {
		imx.RightColumn(c[i].R, c[i].G, c[i].B)
		c[i].R = x // X
		c[i].G = y // Y
		c[i].B = z // Z
	}
	cc := *(*XYZs)(unsafe.Pointer(&c))
	return cc, nil
}

// more flexible method
// you can use your own RGB color space and XYZ white point

func (c XYZs) XYZs() XYZs {
	return c
}

// return self

// DEBUG
func (c XYZs) Show() {
	fmt.Println("CIE XYZ Array, Length %d", len(c))
	for i := 0; i < len(c); i++ {
		fmt.Printf("X : %.48f, Y : %.48f, Z : %.48f\n", c[i].X, c[i].Y, c[i].Z)
	}
}
