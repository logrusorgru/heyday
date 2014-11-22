package heyday

func (c *YXY) XYZ() *XYZ {
	x := c.X * (c.Y1 / c.Y2)
	y := c.Y1
	z := (1 - c.X - c.Y2) * (c.Y1 / c.Y2)
	xyz := &XYZ{x, y, z}
	return xyz
}

func (c *YXY) It() (y1, x, y2 float64) {
	return c.Y1, c.X, c.Y2
}

func (c *YXY) YXY() *YXY {
	return c
}

// Deltas

func (c *YXY) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *YXY) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *YXY) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *YXY) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *YXY) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *YXY) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
