package heyday

func (c *YUV) RGB() *RGB {
	r := c.Y + 1.13983*(c.V-128)
	g := c.Y - 0.39465*(c.U-128) - 0.58060*(c.V-128)
	b := c.Y + 2.03211*(c.U-128)
	rgb := &RGB{r, g, b}
	return rgb
}

func (c *YUV) It() (y, u, v float64) {
	return c.Y, c.U, c.V
}

func (c *YUV) YUV() *YUV {
	return c
}

// Deltas

func (c *YUV) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *YUV) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *YUV) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *YUV) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *YUV) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *YUV) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
