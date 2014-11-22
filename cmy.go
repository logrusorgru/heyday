package heyday

func (c *CMY) RGB() *RGB {
	//CMY values from 0 to 1
	//RGB results from 0 to 255
	r = (1 - c.C) * 255
	g = (1 - c.M) * 255
	b = (1 - c.Y) * 255
	rgb := &RGB{r, g, b}
	return rgb
}

func (c *CMY) CMYK() *CMYK {
	//CMYK and CMY values from 0 to 1
	var c, m, y, k float64
	k = 1

	if c.C < k {
		k = c.C
	}
	if m < k {
		k = c.M
	}
	if y < k {
		k = c.Y
	}
	if k == 1 { //Black
		c = 0
		m = 0
		y = 0
	} else {
		c = (c.C - k) / (1 - k)
		m = (c.M - k) / (1 - k)
		y = (c.Y - k) / (1 - k)
	}
	cmyk := &CMYK{c, m, y, k}
	return cmyk
}

func (c *CMY) It() (cc, m, y float64) {
	return c.C, c.M, c.Y
}

func (c *CMY) CMY() *CMY {
	return c
}

// Deltas

func (c *CMY) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *CMY) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *CMY) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *CMY) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *CMY) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *CMY) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
