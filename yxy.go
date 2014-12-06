package heyday

import (
	. "github.com/logrusorgru/heyday"
)

func (c *Yxy) XYZ() *XYZ {
	x := (c.Y / c.Yc) * c.Xc
	z := (c.Y / c.Yc) * (1 - c.Xc - c.Yc)
	return &XYZ{x, c.Y, z}
}

// ref.: https://en.wikipedia.org/wiki/CIE_1931_color_space#CIE_xy_chromaticity_diagram_and_the_CIE_xyY_color_space
