package heyday

func (c *CMYK) CMY() *CMY {
	//CMYK and CMY values from 0 to 1
	c = (c.C*(1-c.K) + c.K)
	m = (c.M*(1-c.K) + c.K)
	y = (c.Y*(1-c.K) + c.K)
	cmy := &CMY{c, m, y}
	return cmy
}

func (c *CMYK) RGB() *RGB {
	r := 255 * (1 - c.C) * (1 - c.K)
	g := 255 * (1 - c.M) * (1 - c.K)
	b := 255 * (1 - c.Y) * (1 - c.K)
	rgb := &RGB{r, g, b}
	return rgb
}

func (c *CMYK) It() (cc, m, y, k float64) {
	return c.C, c.M, c.Y, c.K
}

func (c *CMYK) CMYK() *CMYK {
	return c
}

// Deltas

func (c *CMYK) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *CMYK) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *CMYK) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *CMYK) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *CMYK) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *CMYK) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
