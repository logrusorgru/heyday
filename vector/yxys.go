package heyday

import (
	. "github.com/logrusorgru/heyday"
)

func (c Yxys) XYZs() XYZs {
	for i := 0; i < len(c); i++ {
		x := (c[i].Y / c[i].Yc) * c[i].Xc
		z := (c[i].Y / c[i].Yc) * (1 - c[i].Xc - c[i].Yc)
		c[i].Y2 = z // Z
		c[i].X = x
	}
	// TRANSFORM the Yxys to XYZs
	cc := *(*XYZs)(unsafe.Pointer(&c))
	return cc
}

// ref.: https://en.wikipedia.org/wiki/CIE_1931_color_space#CIE_xy_chromaticity_diagram_and_the_CIE_xyY_color_space

// DEBUG
func (c Xxys) Show() {
	fmt.Println("CIE Yxy Array, Length %d", len(c))
	for i := 0; i < len(c); i++ {
		fmt.Printf("Y : %.48f, x : %.48f, y : %.48f\n", c[i].Y1, c[i].X, c[i].Y2)
	}
}
