package heyday

import (
	"fmt" // DEBUG
	. "github.com/logrusorgru/heyday"
)

func (c RGBs) XYZs(cio ...int) XYZs {
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
	xyzs, _ := c.XYZsl(
		&rgb_white_points[color_space],
		&white_points[observer][illuminant],
	)
	return xyzs
}

// default color space sRGB
// default illuminant D65
// default observer CIE 1931 2° Observer
// call examples:
//  rgbs.XYZs()                          | sRGB, D65, CIE 1931 2° Observer
//	rgbs.XYZs(Adobe_RGB_1998)            | Adobe_RGB_1998, D65, CIE 1931 2° Observer
//	rgbs.XYZs(Adobe_RGB_1998, D50)       | Adobe_RGB_1998, D50, CIE 1931 2° Observer
//	rgbs.XYZs(Adobe_RGB_1998, D50, O10)  | Adobe_RGB_1998, D50, CIE 1964 10° Observer

func (c RGBs) XYZsl(color_space *Senary, white_point *WP) (XYZs, error) {
	dmx, err := RgbDirectMatrix(
		color_space,
		white_point,
	)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(c); i++ {
		dmx.RightColumn(c[i].R, c[i].G, c[i].B)
		c[i].R = x // X
		c[i].G = y // Y
		c[i].B = z // Z
	}
	cc := *(*XYZs)(unsafe.Pointer(&c))
	return cc, nil
}

// more flexible method
// you can use your own RGB color space and XYZ white point

func (c RGBs) RGBs() RGBs {
	return c
}

// return sefl

// DEBUG
func (c RGBs) Show() {
	fmt.Println("RGB Array, Length %d", len(c))
	for i := 0; i < len(c); i++ {
		fmt.Printf("R : %.48f, G : %.48f, B : %.48f\n", c[i].R, c[i].G, c[i].B)
	}
}
