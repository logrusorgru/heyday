package heyday

import (
	"fmt" // DEBUG
	. "github.com/logrusorgru/heyday"
	"math"
)

func (c Labs) XYZs(io ...int) XYZs {
	var illuminant, observer int
	if len(io) == 2 {
		observer = io[1]
		illuminant = io[0]
	} else if len(io) == 1 {
		illuminant = io[0]
		observer = O2
	} else {
		illuminant = D65
		observer = O2
	}
	return c.XYZsl(&white_points[observer][illuminant])
}

// Create CIE 1931 Color Space (CIE XYZ) from current CIE L*a*b*
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// ref.: ASTM E308 http://wenku.baidu.com/view/1dc90ac20c22590102029dce
//            math http://mathb.in/24271

func (c Labs) XYZsl(wp *WP) XYZs {
	wt := TristimulusWhite(wp.X, wp.Y) // type XYZ
	wp = nil                           // GC (Is it really needed?)
	// main loop
	for i := 0; i < len(c); i++ {
		y := (c[i].L + 16) / 116 // f(Qy)
		x := c[i].A/500 + y      // f(Qx)
		z := y - c[i].B/200      // f(Qz)
		if x > c6d29 {           // eq.: Qx^3 > (6/29)^3 |=> Qx > 6/29
			x = math.Pow(x, 3) // Qx = f(Qx)^3
		} else {
			x = c108d841 * (x - c4d29)
		}
		if y > c6d29 {
			y = math.Pow(y, 3)
		} else {
			y = c108d841 * (y - c4d29)
		}
		if z > c6d29 {
			z = math.Pow(z, 3)
		} else {
			z = c108d841 * (z - c4d29)
		}
		c[i].L = x * wt.X // X = Qx*Xn (<=| Qx = (X/Xn))
		c[i].A = y * wt.Y // 100
		c[i].B = z * wt.Z
	}
	// TRANSFORM Labs to XYZs
	cc := *(*XYZs)(unsafe.Pointer(&c))
	return cc
}

// more flexible method
// you can use your own XYZ white point

func (c Labs) Hue() []float64 {
	hues := []float64{}
	for i := 0; i < len(c); i++ {
		hues = append(hues, hue(c[i].A, c[i].B))
	}
	return hues
}

// return CIE 1976 hue angle [0, 360]° of current color point

func (c Labs) Chromas() []float64 {
	cs := []float64{}
	for i := 0; i < len(c); i++ {
		cs = append(cs, chromas(c[i].A, c[i].B))
	}
	return cs
}

// return CIE 1976 chromas o current color point

func (c Labs) LCHabs() LCHabs {
	for i := 0; i < len(c); i++ {
		h := hue(c.A, c.B)
		cc := chromas(c.A, c.B)
		// an L is tthe same
		c[i].A = cc // C
		c[i].B = h  // H
	}
	// TRANSFORM Labs to XYZs
	cc := *(*LCHabs)(unsafe.Pointer(&c))
	return cc
}

// create L*CH° version of CIE L*,a*,b*

func (c Labs) Labs() Labs {
	return c
}

// return sefl

// DEBUG
func (c Labs) Show() {
	fmt.Println("CIE L*,a*,b* Array, Length %d", len(c))
	for i := 0; i < len(c); i++ {
		fmt.Printf("L*: %.48f, u*: %.48f, v*: %.48f\n", c[i].L, c[i].A, c[i].B)
	}
}
