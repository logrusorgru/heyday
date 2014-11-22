package heyday

func (c *HunterLAB) XYZ() *XYZ {
	y := c.L / 10
	x := c.A / 17.5 * c.L / 10
	z := c.B / 7 * c.L / 10

	y = y ^ 2
	x = (x + y) / 1.02
	z = -(z - y) / 0.847

	xyz := &XYZ{x, y, z}
	return xyz
}

// return self values

func (c *HunterLAB) It() (l, a, b float64) {
	return c.L, c.A, c.B
}

// return self

func (c *HunterLAB) HunterLAB() *HunterLAB {
	return c
}

// Deltas

func (c *HunterLAB) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *HunterLAB) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *HunterLAB) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *HunterLAB) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *HunterLAB) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *HunterLAB) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
