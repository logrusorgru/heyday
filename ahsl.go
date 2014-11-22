package heyday

// https://ru.wikipedia.org/wiki/AHSL

// own

func (c *AHSL) RGB() *RGB {
	var r, g, b float64
	h, s, l := c.H, c.S, c.L
	if h == 0 {
		r = 255
	}
	if h > 0 && h <= 60 {
		r = 255
		g = 255 / (60 / h)
	}
	if h > 60 && h <= 120 {
		g = 255
		r = 255 - 255/(60/(h-60))
	}
	if h > 120 && h <= 180 {
		g = 255
		b = 255 / (60 / (h - 120))
	}
	if h > 180 && h <= 240 {
		b = 255
		g = 255 - 255/(60/(h-180))
	}
	if h > 240 && h <= 300 {
		b = 255
		r = 255 / (60 / (h - 240))
	}
	if h > 300 && h <= 360 {
		r = 255
		b = 255 - 255/(60/(h-300))
	}
	if l > 0 {
		r = r + (255-r)/(100/l)
		g = g + (255-g)/(100/l)
		b = b + (255-b)/(100/l)
	} else if l < 0 {
		r = r - r/(-100/l)
		g = g - g/(-100/l)
		b = b - b/(-100/l)
	}
	gray := (r + g + b) / 3
	if s == 0 {
		r = gray
		g = r
		b = r
	} else if s < 255 {
		r = gray + (r-gray)/(255/s)
		g = gray + (g-gray)/(255/s)
		b = gray + (b-gray)/(255/s)
	}

	rgb := &RGB{r, g, b}
	return rgb
}

func (c *AHLS) It() (h, l, s float64) {
	return c.H, c.L, c.S
}

func (c *AHLS) AHSL() *AHSL {
	return c
}

// Through

func (c *AHSL) RGBA() *RGBA { //
	return c.RGB().RGBA()
}
func (c *AHSL) CMY() *CMY { //
	return c.RGB().CMY()
}
func (c *AHSL) CMYK() *CMYK { //
	return c.RGB().CMYK()
}
func (c *AHSL) HSL() *HSL { //
	return c.RGB().HSL()
}
func (c *AHSL) HSLA() *HSLA { //
	return c.RGB().HSLA()
}
func (c *AHSL) HSB() *HSB { //
	return c.RGB().HSB()
}
func (c *AHSL) HSV() *HSV { //
	return c.RGB().HSV()
}
func (c *AHSL) LAB() *LAB {
	return c.RGB().LAB()
}
func (c *AHSL) HunterLAB() *HunterLAB {
	return c.RGB().HunterLAB()
}
func (c *AHSL) LCH() *LCH {
	return c.RGB().LCH()
}
func (c *AHSL) LUV() *LUV {
	return c.RGB().LUV()
}
func (c *AHSL) XYZ() *XYZ { //
	return c.RGB().XYZ()
}
func (c *AHSL) YXY() *YXY {
	return c.RGB().YXY()
}
func (c *AHSL) YCbCr() *YCbCr { //
	return c.RGB().YCbCr()
}
func (c *AHSL) YIQ() *YIQ { //
	return c.RGB().YIQ()
}
func (c *AHSL) YUV() *YUV { //
	return c.RGB().YUV()
}

// Deltas

func (c *AHSL) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *AHSL) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *AHSL) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *AHSL) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *AHSL) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *AHSL) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
