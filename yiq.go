package heyday

func (c *YIQ) RGB() *RGB {
	r := c.Y + 0.956*c.I + 0.623*c.Q
	g := c.Y - 0.272*c.I - 0.648*c.Q
	b := c.Y - 1.105*c.I + 1.705*c.Q
	rgb := &RGB{r, g, b}
	return rgb
}

// return self values

func (c *YIQ) It() (y, i, q float64) {
	return c.Y, c.I, c.Q
}

// return self

func (c *YIQ) YIQ() *YIQ {
	return c
}

// Deltas

func (c *YIQ) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *YIQ) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *YIQ) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *YIQ) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *YIQ) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *YIQ) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
