package heyday

import "math"

/*
 Observer					2° (CIE 1931)			|		10° (CIE 1964)

 Illuminant			X2			Y2			Z2			X10			Y10			Z10

 A (Incandescent)	109.850		100			35.585		111.144		100			35.200
 C					98.074		100			118.232		97.285		100			116.145
 D50				96.422		100			82.521		96.720		100			81.427
 D55				95.682		100			92.149		95.799		100			90.926
 D65 (Daylight)	  [ 95.047		100			108.883 ]	94.811		100			107.304		default
 D75				94.972		100			122.638		94.416		100			120.641
 F2 (Fluorescent)	99.187		100			67.395		103.280		100			69.026
 F7					95.044		100			108.755		95.792		100			107.687
 F11				100.966		100			64.370		103.866		100			65.627
*/

const (
	A		int = iota
	C
	D50
	D55
	D65     // default
	D75
	F2
	F7
	F11
	Ad10
	Cd10
	D50d10
	D55d10
	D65d10
	D75d10
	F2d10
	F7d10
	F11d10
)

type refxyz {
	x, z float64
}

var refXYZ = map[int]refxyz{}

func init() {
	// ref Y2 and Y10 always is equal to 100
	refXYZ[A]	= refxyz{	109.850,	35.585 	}
	refXYZ[C]	= refxyz{	98.074,		118.232 }
	refXYZ[D50]	= refxyz{	96.422,		82.521 	}
	refXYZ[D55]	= refxyz{	95.682,		92.149 	}
	refXYZ[D65]	= refxyz{	95.047,		108.883 }
	refXYZ[D75]	= refxyz{	94.972,		122.638 }
	refXYZ[F2]	= refxyz{	99.187,		67.395 	}
	refXYZ[F7]	= refxyz{	95.044,		108.755 }
	refXYZ[F11]	= refxyz{	100.966,	64.370 	}
	refXYZ[Ad10]	= refxyz{	111.144,	35.200	}
	refXYZ[Cd10]	= refxyz{	97.285,		116.145	}
	refXYZ[D50d10]	= refxyz{	96.720,		81.427	}
	refXYZ[D55d10]	= refxyz{	95.799,		90.926	}
	refXYZ[D65d10]	= refxyz{	94.811,		107.304	}
	refXYZ[D75d10]	= refxyz{	94.416,		120.641	}
	refXYZ[F2d10]	= refxyz{	103.280,	69.026	}
	refXYZ[F7d10]	= refxyz{	95.792,		107.687	}
	refXYZ[F11d10]	= refxyz{	103.866,	65.627	}
}

// Observer = 2°, Illuminant = D65 by default
func (c *XYZ) RGB() *RGB {
	x := c.X / 100
	y := x.Y / 100
	z := c.Z / 100
	r := x*3.2406 + y*-1.5372 + z*-0.4986
	g := x*-0.9689 + y*1.8758 + z*0.0415
	b := x*0.0557 + y*-0.2040 + z*1.0570
	if r > 0.0031308 {
		r = 1.055*(r^(1/2.4)) - 0.055
	} else {
		r = 12.92 * r
	}
	if g > 0.0031308 {
		g = 1.055*(g^(1/2.4)) - 0.055
	} else {
		g = 12.92 * g
	}
	if b > 0.0031308 {
		b = 1.055*(b^(1/2.4)) - 0.055
	} else {
		b = 12.92 * b
	}
	rgb := &RGB{r * 255, g * 255, b * 255}
	return rgb
}

func (c *XYZ) YXY() *YXY {
	y1 := c.Y
	x := c.X / (c.X + c.Y + c.Z)
	y2 := c.Y / (c.X + c.Y + c.Z)
	yxy := &YXY{y1, x, y2}
	return yxy
}

func (c *XYZ) HunterLAB() *HunterLAB {
	l := 10 * math.Sqrt(c.Y)
	a := 17.5 * (((1.02 * c.X) - c.Y) / math.Sqrt(c.Y))
	b := 7 * ((c.Y - (0.847 * c.Z)) / math.Sqrt(c.Y))
	lab := &HunterLAB{l, a, b}
	return lab
}

// Observer= 2°, Illuminant= D65 by default
// may panic
func (c *XYZ) LAB(oi ...int) *LAB {
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
			panic( errors.New("No such Observer/Illuminant values: %v", oi[0]) )
		}
	}
	x := x.X / refX
	y := c.Y / refY
	z := c.Z / refZ

	if x > 0.008856 {
		x = x ^ (1 / 3)
	} else {
		x = (7.787 * x) + (16 / 116)
	}
	if y > 0.008856 {
		y = y ^ (1 / 3)
	} else {
		y = (7.787 * y) + (16 / 116)
	}
	if z > 0.008856 {
		z = z ^ (1 / 3)
	} else {
		z = (7.787 * z) + (16 / 116)
	}

	l := (116 * y) - 16
	a := 500 * (x - y)
	b := 200 * (y - z)

	lab := &LAB{l, a, b}
	return lab
}

// Observer= 2°, Illuminant= D65 by default
// may panic
func (c *XYZ) LUV(oi ...int) *LUV {
	u := (4 * c.X) / (c.X + (15 * c.Y) + (3 * c.Z))
	v := (9 * c.Y) / (c.X + (15 * c.Y) + (3 * c.Z))

	y := c.Y / 100
	if y > 0.008856 {
		y = y ^ (1 / 3)
	} else {
		y = (7.787 * y) + (16 / 116)
	}

	/*	luv.go	*/
	refU, refV := refUVbyOI(oi...)

	l := (116 * y) - 16
	u = 13 * l * (u - refU)
	v = 13 * l * (varV - refV)

	luv := &LUV{l, u, v}
	return luv
}

// return self values

func (c *XYZ) It() (x, y, z float64) {
	return c.X, c.Y, c.Z
}

// return self

func (c *XYZ) XYZ() *XYZ {
	return c
}

// Deltas

func (c *XYZ) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *XYZ) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *XYZ) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *XYZ) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *XYZ) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *XYZ) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
