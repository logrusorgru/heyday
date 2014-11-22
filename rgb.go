package heyday

import (
	"fmt"
	"math"
)

func (c *RGB) RGBA() *RGBA {
	rgba := &RGBA{c.R, c.G, c.B, 1} // alpha = 1
	return rgba
}

// //Observer. = 2°, Illuminant = D65
func (c *RGB) XYZ() *XYZ {
	r := (c.R / 255) //R from 0 to 255
	g := (c.G / 255) //G from 0 to 255
	b := (c.B / 255) //B from 0 to 255

	if r > 0.04045 {
		r = ((r + 0.055) / 1.055) ^ 2.4
	} else {
		r = r / 12.92
	}
	if g > 0.04045 {
		g = ((g + 0.055) / 1.055) ^ 2.4
	} else {
		g = g / 12.92
	}
	if b > 0.04045 {
		b = ((b + 0.055) / 1.055) ^ 2.4
	} else {
		b = b / 12.92
	}

	r = r * 100
	g = g * 100
	b = b * 100

	x := r*0.4124 + g*0.3576 + b*0.1805
	y := r*0.2126 + g*0.7152 + b*0.0722
	z := r*0.0193 + g*0.1192 + b*0.9505

	xyz := &XYZ{x, y, z}
	return xyz
}

// lost quality
// http://play.golang.org/p/9q5yBNDh3W

func (c *RGB) HSL() *HSL {
	r := (c.R / 255) //RGB from 0 to 255
	g := (c.G / 255)
	b := (c.B / 255)

	min := math.Min(r, math.Min(g, b)) //Min. value of RGB
	max := math.Max(r, math.Max(g, b)) //Max. value of RGB
	del := max - min                   //Delta RGB value

	l := (max + min) / 2

	var h, s float64

	if del == 0 { //This is a gray, no chroma...
		h = 0 //HSL results from 0 to 1
		s = 0
	} else { //Chromatic data...
		if l < 0.5 {
			s = del / (max + min)
		} else {
			s = del / (2 - max - min)
		}

		delR := (((max - r) / 6) + (del / 2)) / del
		delG := (((max - g) / 6) + (del / 2)) / del
		delB := (((max - b) / 6) + (del / 2)) / del

		switch max {
		case r:
			h = delB - delG
		case g:
			h = (1 / 3) + delR - delB
		case b:
			h = (2 / 3) + delG - delR
		}

		if h < 0 {
			h += 1
		}
		if h > 1 {
			h -= 1
		}
	}
	hsl := &HSL{h, s, l}
	return hsl
}

func (c *RGB) HSLA() *HSLA {
	t := c.HSL()
	hsla := &HSLA{t.H, t.S, t.L, 1} // alpha = 1
	return hsla
}

func (c *RGB) AHSL() *AHSL {
	var h, s, l float64
	r, g, b := c.R, c.G, c.B
	gray := (r + g + b) / 3
	min := math.Min(r, math.Min(g, b)) //Min. value of RGB
	max := math.Max(r, math.Max(g, b)) //Max. value of RGB
	del := max - min                   //Delta RGB value
	if del == 0 {
		h = 0
	} else {
		if max == r && g >= b {
			h = 60 * (g - b) / del
		}
		if max == r && b > g {
			h = 60*(g-b)/del + 360
		}
		if max == g {
			h = 60*(b-r)/del + 120
		}
		if max == b {
			h = 60*(r-g)/del + 240
		}
	}
	var r0, g0, b0 float64 = 0, 0, 0
	if h == 0 {
		r0 = 255
	}
	if h > 0 && h <= 60 {
		r0 = 255
		g0 = 255 / (60 / h)
	}
	if h > 60 && h <= 120 {
		g0 = 255
		r0 = 255 - 255/(60/(h-60))
	}
	if h > 120 && h <= 180 {
		g0 = 255
		b0 = 255 / (60 / (h - 120))
	}
	if h > 180 && h <= 240 {
		b0 = 255
		g0 = 255 - 255/(60/(h-180))
	}
	if h > 240 && h <= 300 {
		b0 = 255
		r0 = 255 / (60 / (h - 240))
	}
	if h > 300 && h <= 360 {
		r0 = 255
		b0 = 255 - 255/(60/(h-300))
	}
	gray0 := (r0 + g0 + b0) / 3
	if gray > gray0 {
		l = 100 / ((255 - gray0) / (gray - gray0))
	}
	if gray < gray0 {
		l = 100 / ((gray0) / (gray - gray0))
	}
	if gray == gray0 {
		l = 0
	}
	if l > 0 {
		r0 = r0 + (255-r0)/(100/l)
		g0 = g0 + (255-g0)/(100/l)
		b0 = b0 + (255-b0)/(100/l)
	} else if l < 0 {
		r0 = r0 - r0/(-100/l)
		g0 = g0 - g0/(-100/l)
		b0 = b0 - b0/(-100/l)
	}
	if math.Abs(r-gray) == 0 {
		s = 0
	} else {
		s = 255 / (math.Abs(r0-gray) / math.Abs(r-gray))
	}
	hsl := &AHSL{h, s, l}
	return hsl
}

// lost quality
// http://play.golang.org/p/9q5yBNDh3W

func (c *RGB) HSV() *HSV {
	r := (c.R / 255) //RGB from 0 to 255
	g := (c.G / 255)
	b := (c.B / 255)

	min := math.Min(r, math.Min(g, b)) //Min. value of RGB
	max := math.Max(r, math.Max(g, b)) //Max. value of RGB
	del := max - min                   //Delta RGB value

	v := max

	var h, s float64

	if del == 0 { //This is a gray, no chroma...
		h = 0 //HSV results from 0 to 1
		s = 0
	} else { //Chromatic data...
		s = del / max

		delR := (((max - r) / 6) + (del / 2)) / del
		delG := (((max - g) / 6) + (del / 2)) / del
		delB := (((max - b) / 6) + (del / 2)) / del

		switch max {
		case r:
			h = delB - delG
		case g:
			h = (1 / 3) + delR - delB
		case b:
			h = (2 / 3) + delG - delR
		}

		if h < 0 {
			h += 1
		}
		if h > 1 {
			h -= 1
		}
	}
	hsv := &HSV{h, s, v}
	return hsv
}

func (c *RGB) HSB() *HSB {
	hsv := c.HSV()
	hsb := &HSB{hsv.H, hsv.S, hsv.V}
	return hsb
}

func (c *RGB) CMY() *CMY {
	//RGB values from 0 to 255
	//CMY results from 0 to 1
	c := 1 - (c.R / 255)
	m := 1 - (c.G / 255)
	y := 1 - (c.B / 255)
	cmy := &CMY{c, m, y}
	return cmy
}

func (c *RGB) CMYK() *CMYK {
	//RGB values from 0 to 255
	//CMY results from 0 to 1
	r := c.R / 255
	g := c.G / 255
	b := c.B / 255
	k := 1 - math.Max(r, math.Max(g, b))
	c := (1 - r - k) / (1 - k)
	m := (1 - g - k) / (1 - k)
	y := (1 - b - k) / (1 - k)
	cmyk := &CMYK{c, m, y, k}
	return cmyk
}

// RGB to Y'CbCr
func (c *RGB) YCbCr() *YCbCr {
	// https://code.google.com/p/go/source/browse/src/pkg/image/color/ycbcr.go?name=release#8
	// The JFIF specification says:
	//      Y' =  0.2990*R + 0.5870*G + 0.1140*B
	//      Cb = -0.1687*R - 0.3313*G + 0.5000*B + 128
	//      Cr =  0.5000*R - 0.4187*G - 0.0813*B + 128
	// http://www.w3.org/Graphics/JPEG/jfif3.pdf says Y but means Y'.
	// fixed with: https://ru.wikipedia.org/wiki/YCbCr#JPEG_.D0.BF.D1.80.D0.B5.D0.BE.D0.B1.D1.80.D0.B0.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D1.8F
	y := 0.299*c.R + 0.587*c.G + 0.114*c.B
	cb := -0.168736*c.R - 0.331264*c.G + 0.5*c.B + 128
	cr := 0.5*c.R - 0.418688*c.G - 0.081312*c.B + 128
	if y < 0 {
		y = 0
	} else if y > 255 {
		y = 255
	}
	if cb < 0 {
		cb = 0
	} else if cb > 255 {
		cb = 255
	}
	if cr < 0 {
		cr = 0
	} else if cr > 255 {
		cr = 255
	}
	ycbcr := &YCbCr{y, cb, cr}
	return ycbcr
}

func (c *RGB) YIQ() *YIQ {
	y := 0.299*c.R + 0.587*c.G + 0.114*c.B
	i := 0.596*c.R - 0.274*c.G - 0.322*c.B
	q := 0.211*c.R - 0.522*c.G + 0.311*c.B
	yiq := &YIQ{y, i, q}
	return yiq
}

func (c *RGB) YUV() *YUV {
	y := 0.299*c.R + 0.587*c.G + 0.114*c.B
	u := -0.14713*c.R - 0.28886*c.G + 0.436*c.B + 128
	v := 0.615*c.R - 0.51499*c.G - 0.10001*c.B + 128
	yuv := &YUV{y, u, v}
	return yuv
}

// Through

func (c *RGB) LAB() *LAB { //
	return c.XYZ().LAB()
}
func (c *RGB) HunterLAB() *HunterLAB { //
	return c.XYZ().HunterLAB()
}
func (c *RGB) LCH() *LCH {
	return c.XYZ().LCH()
}
func (c *RGB) LUV() *LUV { //
	return c.XYZ().LUV()
}
func (c *RGB) YXY() *YXY { //
	return c.XYZ().YXY()
}

// Other

func (c *RGB) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", round(c.R), round(c.G), round(c.B))
}

func (c *RGB) Short() string {
	return fmt.Sprintf(
		"#%x%x%x",
		round(c.R/17)*17/16,
		round(c.G/17)*17/16,
		round(c.B/17)*17/16,
	)
}

func (c *RGB) Websafe() string {
	return fmt.Sprintf(
		"#%02x%02x%02x",
		round(c.R/51)*51,
		round(c.G/51)*51,
		round(c.B/51)*51,
	)
}

func (c *RGB) ShortWebsafe() string {
	return fmt.Sprintf(
		"#%x%x%x",
		round(c.R/51)*51/16,
		round(c.G/51)*51/16,
		round(c.B/51)*51/16,
	)
}

// only this type (RGB) have this method
// return string to css, like `rgb(128, 200, 50)`
func (c *RGB) Css() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", round(c.R), round(c.G), round(c.B))
}

// append alpha channel, create new RGBA color

func (c *RGB) AddAlpha(a float64) *RGBA {
	rgba := &RGBA{
		c.R,
		c.G,
		c.B,
		a,
	}
	return rgba
}

// retunr self values

func (c *RGB) It() (r, g, b float64) {
	return c.R, c.G, c.B
}

// return self

func (c *RGB) RGB() *RGB {
	return c
}

// Deltas

func (c *RGB) DeltaC(d *Color) float64 {
	return c.LAB().DeltaC(d)
}
func (c *RGB) DeltaH(d *Color) float64 {
	return c.LAB().DeltaH(d)
}
func (c *RGB) DeltaE(d *Color) float64 {
	return c.LAB().DeltaE(d)
}
func (c *RGB) DeltaE94(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE94(d, w...)
}
func (c *RGB) DeltaE00(d *Color, w ...float64) float64 {
	return c.LAB().DeltaE00(d, w...)
}
func (c *RGB) DeltaCMC(d *Color, w ...float64) float64 {
	return c.LAB().DeltaCMC(d, w...)
}
