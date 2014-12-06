package heyday

import (
	"fmt" // DEBUG
	. "github.com/logrusorgru/heyday"
	"math"
)

//
// CIE 1976 (L*, u*, v*) color space (??? Is it true? ???)
//

func (c Luvs) XYZs(io ...int) XYZs {
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

// Transform to CIE 1931 Color Space (CIE XYZ) from current CIE L*,u*,v*
// By default used CIE 1931 2° Observer and D65 Standart Illuminant.
// But you can also choose particular illuminant and observer.
// First param  - set illuminant (optional)
// Second param - set observer (optional)
// ref.: ASTM E308 http://wenku.baidu.com/view/1dc90ac20c22590102029dce

func (c Luvs) XYZsl(wp *WP) XYZs {
	wt := TristimulusWhite(wp.X, wp.Y)        // type XYZ
	wp = nil                                  // GC (Is it really needed?)
	usn := 4 * wt.X / (wt.X + c1500 + 3*wt.Z) // uₙ'
	vsn := 9 * wt.Y / (wt.X + c1500 + 3*wt.Z) // vₙ'
	// main loop
	for i := 0; i < len(c); i++ {
		y := (c[i].L + 16) / 116 // pre-f(Qy)
		if y > c6d29cbrt {
			y = math.Pow(y, 3) // Qy
		} else {
			y = c108d841 * (y - c4d29) // Qy
		}
		us := c[i].U/(13*c[i].L) + usn
		vs := c[i].V/(13*c[i].L) + vsn
		y = y * wt.Y                                 // Qy = Y/Yₙ |=> Y = Qy * Yₙ
		c[i].L = -(9 * y * us) / ((us-4)*vs - us*vs) // X
		c[i].U = y                                   // Y
		c[i].V = (9*y - 15*vs*y - vs*x) / (3 * vs)   // Z
	}
	// TRANSFORM Labs to XYZs
	cc := *(*XYZs)(unsafe.Pointer(&c))
	return cc
}

// more flexible method
// you can use your own XYZ white point

func (c Luvs) Hue() []float64 {
	hues := []float64{}
	for i := 0; i < len(c); i++ {
		hues = append(hues, hue(c[i].U, c[i].V))
	}
	return hues
}

// return CIE 1976 hue angle [0, 360]° of current color point

func (c Luvs) Chromas() []float64 {
	cs := []float64{}
	for i := 0; i < len(c); i++ {
		cs = append(cs, chromas(c[i].U, c[i].V))
	}
	return cs
}

// return CIE 1976 chromas o current color point

func (c Luvs) LCHuvs() LCHuvs {
	for i := 0; i < len(c); i++ {
		h := hue(c[i].U, c[i].V)
		cc := chromas(c[i].U, c[i].V)
		// an L is the same
		c[i].U = cc // C
		c[i].V = h  // H
	}
	// TRANSFORM Labs to XYZs
	cc := *(*LCHuvs)(unsafe.Pointer(&c))
	return cc
}

// create L*CH°(uv) version of CIE L*,a*,b*

func (c Luvs) Luvs() Luvs {
	return c
}

// return self

// DEBUG
func (c Luvs) Show() {
	fmt.Println("CIE L*,u*,v* Array, Length %d", len(c))
	for i := 0; i < len(c); i++ {
		fmt.Printf("L*: %.48f, u*: %.48f, v*: %.48f\n", c[i].L, c[i].U, c[i].V)
	}
}
