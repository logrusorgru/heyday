package heyday

func (c *YcbCr) RGB() *RGB {
	// https://code.google.com/p/go/source/browse/src/pkg/image/color/ycbcr.go?name=release#8
	// The JFIF specification says:
	//      R = Y' + 1.40200*(Cr-128)
	//      G = Y' - 0.34414*(Cb-128) - 0.71414*(Cr-128)
	//      B = Y' + 1.77200*(Cb-128)
	// http://www.w3.org/Graphics/JPEG/jfif3.pdf says Y but means Y'.
	// and: https://ru.wikipedia.org/wiki/YCbCr#JPEG_.D0.BF.D1.80.D0.B5.D0.BE.D0.B1.D1.80.D0.B0.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D1.8F
	r := c.Y + 1.40200*(c.Cr-128)
	g := c.Y - 0.34414*(c.Cb-128) - 0.71414*(c.Cr-128)
	b := c.Y + 1.77200*(c.Cb-128)
	if r < 0 {
		r = 0
	} else if r > 255 {
		r = 255
	}
	if g < 0 {
		g = 0
	} else if g > 255 {
		g = 255
	}
	if b < 0 {
		b = 0
	} else if b > 255 {
		b = 255
	}
	rgb := &RGB{r, g, b}
	return rgb
}

// return self values

func (c *YCbCr) It() (y, cb, cr float64) {
	return c.Y, c.Cb, c.Cr
}

// return self

func (c *YCbCr) YCbCr() *YCbCr {
	return c
}

// Deltas

func (c *YCbCr) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *YCbCr) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *YCbCr) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *YCbCr) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *YCbCr) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *YCbCr) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
