package heyday

func (c *XYZ) Chromaticity() (x, y float64) {
	x = c.X / (c.X + c.Y + c.Z)
	y = c.Y / (c.X + c.Y + c.Z)
	// z = 1 - x - y
	return
}

// get chromaticity co-ordinates x, y of current color
