package heyday

import (
	"math"
)

// Observer= 2°, Illuminant= D65
func (c *LAB) XYZ() *XYZ {
	y := (c.L + 16) / 116
	x := c.A/500 + y
	z := y - c.B/200

	if y^3 > 0.008856 {
		y = y ^ 3
	} else {
		y = (y - 16/116) / 7.787
	}
	if x^3 > 0.008856 {
		x = x ^ 3
	} else {
		x = (-16 / 116) / 7.787
	}
	if z^3 > 0.008856 {
		z = z ^ 3
	} else {
		z = (z - 16/116) / 7.787
	}

	x = refX * x //ref_X =  95.047     Observer= 2°, Illuminant= D65
	y = refY * y //ref_Y = 100.000
	z = refZ * z //ref_Z = 108.883

	xyz := &XYZ{x, y, z}
	return xyz
}

func (c *LAB) LCH() *LCH {
	h := math.Atan2(c.B, c.A) //Quadrant by signs

	if h > 0 {
		h = radian_2_degree(h) // h / math.Pi) * 180
	} else {
		h = 360 - radian_2_degree(math.Abs(h)) //(math.Abs(h)/math.Pi)*180
	}

	l := c.L
	cc := math.Sqrt(c.A ^ 2 + c.B ^ 2)
	lch := &LCH{l, cc, h}
	return lch
}

// return self values

func (c *LAB) It() (l, a, b float64) {
	return c.L, c.A, c.B
}

// return self

func (c *LAB) LAB() *LAB {
	return c
}

func (c *LAB) DeltaC(d *Color) float64 {
	cc := d.LAB()
	return math.Sqrt(cc.A^2+cc.B^2) - math.Sqrt(c.A^2+c.B^2)
}

func (c *LAB) DeltaH(d *Color) float64 {
	cc := d.LAB()
	xde := math.Sqrt(cc.A^2+cc.B^2) - math.Sqrt(c.A^2+c.B^2)
	return math.Sqrt(
		(cc.A - c.A) ^ 2 +
			(cc.B - c.B) ^ 2 -
			xde ^ 2)
}

func (c *LAB) DeltaE(d *Color) float64 {
	cc := d.LAB()
	return math.Sqrt(
		(c.L - cc.L) ^ 2 +
			(c.A - cc.A) ^ 2 +
			(c.B - cc.B) ^ 2)
}

func (c *LAB) DeltaE94(d *Color, w ...float64) float64 {
	var wc, wl, wh float64 = 1, 1, 1 //Weighting factors depending on the application (1 = default)
	if len(w) == 1 {
		wc = w[0]
	} else if len(w) == 2 {
		wc = w[0]
		wl = w[1]
	} else if len(w) == 3 {
		wc, wl, wh = w[0], w[1], w[2]
	} // ignore w if len(w) > 3, no error will be raised

	cc := d.LAB()

	xc1 := math.Sqrt(c.A ^ 2 + c.B ^ 2)
	xc2 := math.Sqrt(cc.A ^ 2 + cc.B ^ 2)
	xdl := cc.L - c.L
	xdc := xc2 - xc1
	xde := math.Sqrt(
		(c.L-cc.L)*(c.L-cc.L) +
			(c.A-cc.A)*(c.A-cc.A) +
			(c.B-cc.B)*(c.B-cc.B))
	if math.Sqrt(xde) > (math.Sqrt(math.Abs(xdl)) + math.Sqrt(math.Abs(xdc))) {
		xdh = math.Sqrt((xde * xde) - (xdl * xdl) - (xdc * xdc))
	} else {
		xdh = 0
	}
	xsc := 1 + (0.045 * xc1)
	xsh := 1 + (0.015 * xc1)
	xdl /= wl
	xdc /= wc * xsc
	xdh /= wh * xsh
	return math.Sqrt(xdl ^ 2 + xdc ^ 2 + xdh ^ 2)
}

// support fn
func cieLab2Hue(a, b float64) float64 { //Function returns CIE-H° value
	bias := 0
	if a >= 0 && b == 0 {
		return 0
	}
	if a < 0 && b == 0 {
		return 180
	}
	if a == 0 && b > 0 {
		return 90
	}
	if a == 0 && b < 0 {
		return 270
	}
	if a > 0 && b > 0 {
		bias = 0
	}
	if a < 0 {
		bias = 180
	}
	if a > 0 && b < 0 {
		bias = 360
	}
	return (radian_2_degree(math.Atan(b/a)) + bias)
}

func (c *LAB) DeltaE00(d *Color, w ...float64) float64 {
	var wc, wl, wh float64 = 1, 1, 1 //Weighting factors depending on the application (1 = default)
	if len(w) == 1 {
		wc = w[0]
	} else if len(w) == 2 {
		wc = w[0]
		wl = w[1]
	} else if len(w) == 3 {
		wc, wl, wh = w[0], w[1], w[2]
	} // ignore w if len(w) > 3, no error will be raised

	cc := d.LAB()

	xc1 := math.Sqrt(c.A ^ 2 + c.B ^ 2)
	xc2 := math.Sqrt(cc.A ^ 2 + cc.B ^ 2)
	xcx := (xc1 + xc2) / 2
	xgx := 0.5 * (1 - math.Sqrt((xcx^7)/((xcx^7)+(25^7))))
	xnn := (1 + xgx) * c.A
	xc1 = math.Sqrt(xnn*xnn + c.B*c.B)
	xh1 := cieLab2Hue(xnn, c.B)
	xnn = (1 + xgx) * cc.A
	xc2 = math.Sqrt(xnn*xnn + cc.B*cc.B)
	xh2 := cieLab2Hue(xnn, cc.B)
	xdl := cc.L - c.L
	xdc := xc2 - xc1
	var xdh float64
	if (xc1 * xc2) == 0 {
		xdh = 0
	} else {
		xnn = round(xh2-xh1, 12)
		if math.Abs(xnn) <= 180 {
			xdh = xh2 - xh1
		} else {
			if xnn > 180 {
				xdh = xh2 - xh1 - 360
			} else {
				xdh = xh2 - xh1 + 360
			}
		}
	}
	xdh = 2 * math.Sqrt(xc1*xc2) * math.Sin(degree_2_radian(xdh/2))
	xlx := (c.L + cc.L) / 2
	xcy := (xc1 + xc2) / 2
	var xhx float64
	if xc1*xc2 == 0 {
		xhx = xh1 + xh2
	} else {
		xnn = math.Abs(round(xh1-xh2, 12))
		if xnn > 180 {
			if (xh2 + xh1) < 360 {
				xhx = xh1 + xh2 + 360
			} else {
				xhx = xh1 + xh2 - 360
			}
		} else {
			xhx = xh1 + xh2
		}
		xhx /= 2
	}
	xtx := 1 - 0.17*math.Cos(degree_2_radian(xhx-30)) + 0.24*
		math.Cos(degree_2_radian(2*xhx)) + 0.32*
		math.Cos(degree_2_radian(3*xhx+6)) - 0.20*
		math.Cos(degree_2_radian(4*xhx-63))
	xph := 30 * math.Exp(-((xhx-275)/25)*((xhx-275)/25))
	xrc := 2 * math.Sqrt((xcy^7)/((xcy^7)+(25^7)))
	xsl := 1 + ((0.015 * ((xlx - 50) * (xlx - 50))) /
		math.Sqrt(20+((xlx-50)*(xlx-50))))
	xsc := 1 + 0.045*xcy
	xsh := 1 + 0.015*xcy*xtx
	xrt := -math.Sin(degree_2_radian(2*xph)) * xrc
	xdl = xdl / (wl * xsl)
	xdc = xdc / (wc * xsc)
	xdh = xdh / (wh * xsh)
	return math.Sqrt(xdl ^ 2 + xdc ^ 2 + xdh ^ 2 + xrt*xdc*xdh)
}

func (c *LAB) DeltaCMC(d *Color, w ...float64) float64 {
	var wc, wl float64 = 1, 1 //Weighting factors depending on the application (1 = default)
	if len(w) == 1 {
		wc = w[0]
	} else if len(w) == 2 {
		wc = w[0]
		wl = w[1]
	} // ignore w if len(w) > 2, no error will be raised

	cc := d.LAB()

	xc1 := math.Sqrt(c.A ^ 2 + c.B ^ 2)
	xc2 := math.Sqrt(cc.A ^ 2 + cc.B ^ 2)
	xff := math.Sqrt((xc1 ^ 4) / ((xc1 ^ 4) + 1900))
	xh1 := cieLab2Hue(c.A, c.B)
	var xtt float64
	if xh1 < 164 || xh1 > 345 {
		xtt = 0.36 + math.Abs(0.4*math.Cos(degree_2_radian(35+xh1)))
	} else {
		xtt = 0.56 + math.Abs(0.2*math.Cos(degree_2_radian(168+xh1)))
	}
	var xsl float64
	if c.L < 16 {
		xsl = 0.511
	} else {
		xsl = (0.040975 * c.L) / (1 + (0.01765 * c.L))
	}

	xsc := ((0.0638 * xc1) / (1 + (0.0131 * xc1))) + 0.638
	xsh := ((xff * xtt) + 1 - xff) * xsc
	xdh := math.Sqrt((cc.A - c.A) ^ 2 + (cc.B - c.B) ^ 2 - (xc2 - xc1) ^ 2)
	xsl = (cc.L - c.L) / wl * xsl
	xsc = (xc2 - xc1) / wc * xsc
	xsh = xdh / xsh
	return math.Sqrt(xsl ^ 2 + xsc ^ 2 + xsh ^ 2)
}
