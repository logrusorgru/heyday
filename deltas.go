package heyday

import (
	"math"
)

func (c *Lab) DeltaC(d Color) float64 {
	cc := d.Lab()
	return math.Sqrt(math.Pow(cc.A, 2)+math.Pow(cc.B, 2)) - math.Sqrt(math.Pow(c.A, 2)+math.Pow(c.B, 2))
}

// return Delta C* of current color and 1st param
// if 1st param is not CIE L*a*b* it will be automaticaly converted
// to CIE L*a*b* through CIE XYZ with a little lost of quality
// ref.: http://easyrgb.com/index.php?X=DELT&H=01#text1

func (c *Lab) DeltaH(d Color) float64 {
	cc := d.Lab()
	xde := math.Sqrt(math.Pow(cc.A, 2)+math.Pow(cc.B, 2)) - math.Sqrt(math.Pow(c.A, 2)+math.Pow(c.B, 2))
	return math.Sqrt(
		math.Pow((cc.A-c.A), 2) +
			math.Pow((cc.B-c.B), 2) -
			math.Pow(xde, 2))
}

// return Delta H* of current color and 1st param
// ref.: http://easyrgb.com/index.php?X=DELT&H=02#text2

func (c *Lab) DeltaE(d Color) float64 {
	cc := d.Lab()
	return math.Sqrt(
		math.Pow((c.L-cc.L), 2) +
			math.Pow((c.A-cc.A), 2) +
			math.Pow((c.B-cc.B), 2))
}

// ref.: http://easyrgb.com/index.php?X=DELT&H=03#text3

func (c *Lab) DeltaE94(d Color, w ...float64) float64 {
	var wc, wl, wh float64 = 1, 1, 1 //Weighting factors depending on the application (1 = default)
	if len(w) == 1 {
		wc = w[0]
	} else if len(w) == 2 {
		wc = w[0]
		wl = w[1]
	} else if len(w) == 3 {
		wc, wl, wh = w[0], w[1], w[2]
	} // ignore w if len(w) > 3, no error will be raised

	cc := d.Lab()

	xc1 := math.Sqrt(math.Pow(c.A, 2) + math.Pow(c.B, 2))
	xc2 := math.Sqrt(math.Pow(cc.A, 2) + math.Pow(cc.B, 2))
	xdl := cc.L - c.L
	xdc := xc2 - xc1
	xde := math.Sqrt(
		(c.L-cc.L)*(c.L-cc.L) +
			(c.A-cc.A)*(c.A-cc.A) +
			(c.B-cc.B)*(c.B-cc.B))
	var xdh float64
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
	return math.Sqrt(math.Pow(xdl, 2) + math.Pow(xdc, 2) + math.Pow(xdh, 2))
}

// ref.: http://easyrgb.com/index.php?X=DELT&H=03#text4

func (c *Lab) DeltaE00(d Color, w ...float64) float64 {
	var wc, wl, wh float64 = 1, 1, 1 //Weighting factors depending on the application (1 = default)
	if len(w) == 1 {
		wc = w[0]
	} else if len(w) == 2 {
		wc = w[0]
		wl = w[1]
	} else if len(w) == 3 {
		wc, wl, wh = w[0], w[1], w[2]
	} // ignore w if len(w) > 3, no error will be raised

	cc := d.Lab()

	xc1 := math.Sqrt(math.Pow(c.A, 2) + math.Pow(c.B, 2))
	xc2 := math.Sqrt(math.Pow(cc.A, 2) + math.Pow(cc.B, 2))
	xcx := (xc1 + xc2) / 2.0
	xcxp7 := math.Pow(xcx, 7)
	xgx := 0.5 * (1.0 - math.Sqrt(xcxp7/(xcxp7+c25p7)))
	xnn := (1.0 + xgx) * c.A
	xc1 = math.Sqrt(xnn*xnn + c.B*c.B)
	xh1 := hue(xnn, c.B)
	xnn = (1.0 + xgx) * cc.A
	xc2 = math.Sqrt(xnn*xnn + cc.B*cc.B)
	xh2 := hue(xnn, cc.B)
	xdl := cc.L - c.L
	xdc := xc2 - xc1
	var xdh float64
	if (xc1 * xc2) == 0 {
		xdh = 0
	} else {
		xnn = fround12(xh2 - xh1)
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
		xnn = math.Abs(fround12(xh1 - xh2))
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
	xtx := 1.0 - 0.17*math.Cos(degree_2_radian(xhx-30)) + 0.24*
		math.Cos(degree_2_radian(2*xhx)) + 0.32*
		math.Cos(degree_2_radian(3*xhx+6)) - 0.20*
		math.Cos(degree_2_radian(4*xhx-63))
	xph := 30 * math.Exp(-((xhx-275)/25)*((xhx-275)/25))
	xcyp7 := math.Pow(xcy, 7)
	xrc := 2 * math.Sqrt(xcyp7/(xcyp7+c25p7))
	xsl := 1.0 + ((0.015 * ((xlx - 50) * (xlx - 50))) /
		math.Sqrt(20+((xlx-50)*(xlx-50))))
	xsc := 1.0 + 0.045*xcy
	xsh := 1.0 + 0.015*xcy*xtx
	xrt := -math.Sin(degree_2_radian(2*xph)) * xrc
	xdl = xdl / (wl * xsl)
	xdc = xdc / (wc * xsc)
	xdh = xdh / (wh * xsh)
	return math.Sqrt(math.Pow(xdl, 2) + math.Pow(xdc, 2) + math.Pow(xdh, 2) + xrt*xdc*xdh)
}

// ref.: http://easyrgb.com/index.php?X=DELT&H=03#text5

func (c *Lab) DeltaCMC(d Color, w ...float64) float64 {
	var wc, wl float64 = 1, 1 //Weighting factors depending on the application (1 = default)
	if len(w) == 1 {
		wc = w[0]
	} else if len(w) == 2 {
		wc = w[0]
		wl = w[1]
	} // ignore w if len(w) > 2, no error will be raised

	cc := d.Lab()

	xc1 := math.Sqrt(math.Pow(c.A, 2) + math.Pow(c.B, 2))
	xc2 := math.Sqrt(math.Pow(cc.A, 2) + math.Pow(cc.B, 2))
	xff := math.Sqrt(math.Pow(xc1, 4) / (math.Pow(xc1, 4) + 1900))
	xh1 := hue(c.A, c.B)
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
	xdh := math.Sqrt(math.Pow(cc.A-c.A, 2) + math.Pow(cc.B-c.B, 2) - math.Pow(xc2-xc1, 2))
	xsl = (cc.L - c.L) / wl * xsl
	xsc = (xc2 - xc1) / wc * xsc
	xsh = xdh / xsh
	return math.Sqrt(math.Pow(xsl, 2) + math.Pow(xsc, 2) + math.Pow(xsh, 2))
}

// ref.: http://easyrgb.com/index.php?X=DELT&H=03#text6
