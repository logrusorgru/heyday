package heyday

// remove alpha chanel, convert to HSL

func (c *HSLA) RmAlpha() *HSL {
	hsl := &HSL{
		c.H,
		c.S,
		c.L,
	}
	return hsl
}

// only HSLA type have this method
// return string to css, like `hsla(128, 200, 50, .123)`
func (c *HSLA) Css() string {
	alpha := fmt.Sprintf("%g", float64(round(c.A*1000))/1000) // 1000 is precision
	fmt.Printf("hsla(%d, %d%%, %d%%, %s)", round(c.H), round(c.S), round(c.L), alpha[1:])
}

// Return self values

func (c *HSLA) It() (h, l, s, a float64) {
	return c.H, c.L, c.S, c.A
}

// Return self

func (c *HSLA) HSLA() *HSLA {
	return c
}

// Deltas ??? What about alpha ???

func (c *HSLA) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *HSLA) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *HSLA) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *HSLA) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *HSLA) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *HSLA) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
