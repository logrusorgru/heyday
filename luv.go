package heyday

import (
	"math"
)

/* xyz.go
ref_X =  95.047      //Observer= 2°, Illuminant= D65
ref_Y = 100.000
ref_Z = 108.883
*/

//                          refU      refV
func refUVbyOI(oi ...int) (float64, float64) {
	// select observer/illuminant or set default
	var refX, refY, refZ float64
	refY = 100
	if len(oi) == 0 {
		refX = refXYZ[D65].x
		refZ = refXYZ[D65].z
	} else {
		// there are no check of oi length
		if refxz, ok := refXYZ[oi[0]]; ok {
			refX = refxz.x
			refZ = refxz.z
		} else {
			// panic
			panic(errors.New("No such Observer/Illuminant values: %v", oi[0]))
		}
	}
	refU := (4 * refX) / (refX + (15 * refY) + (3 * refZ))
	refV := (9 * refY) / (refX + (15 * refY) + (3 * refZ))
	return refU, refV
}

// Observer= 2°, Illuminant= D65
func (c *LUV) XYZ(oi ...int) *XYZ {
	y := (c.L + 16) / 116
	if y^3 > 0.008856 {
		y = y ^ 3
	} else {
		y = (y - 16/116) / 7.787
	}

	refU, refV := refUVbyOI(oi...)

	u := c.U/(13*c.L) + refU
	v := c.V/(13*c.L) + refV

	y = y * 100
	x := -(9 * y * u) / ((u-4)*v - u*v)
	z := (9*y - (15 * v * y) - (v * x)) / (3 * v)

	xyz := &XYZ{x, y, z}
	return xyz
}

func (c *LUV) LCH() *LCH {
	l := c.L
	cc := math.Sqrt(c.U ^ 2 + c.V ^ 2)
	h := radian_2_degree(math.Atan2(c.V, c.U))
	if h < 0 {
		h = h + 360
	} else if h >= 360 {
		h = h - 360
	}
	lch := &LCH{l, c, h}
	return lch
}

// return self values

func (c *LUV) It() (l, u, v float64) {
	return c.L, c.U, c.V
}

// return self

func (c *LUV) LUV() *LUV {
	return c
}

// Deltas

func (c *LUV) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *LUV) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *LUV) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *LUV) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *LUV) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *LUV) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
