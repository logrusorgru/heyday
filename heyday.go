package heyday

import (
	"errors"
	"math"
	"strconv"
)

// helpers

func degree_2_radian(d float64) float64 {
	return (d / 180) * math.Pi
}

func radian_2_degree(r float64) float64 {
	return (r / math.Pi) * 180
}

func round(f float64) int {
	return int(f + 0.5)
}

// go

type Color interface {
	RGB() *RGB
	RGBA() *RGBA
	CMY() *CMY
	CMYK() *CMYK
	HSL() *HSL
	HSLA() *HSLA
	AHSL() *AHSL
	HSB() *HSB
	HSV() *HSV
	LAB() *LAB
	HunterLAB() *HunterLAB
	LCH() *LCH
	LUV() *LUV
	XYZ() *XYZ
	YXY() *YXY
	YCbCr() *YCbCr
	YIQ() *YIQ
	YUV() *YUV
	// Return values of color, such as RGB => R, G, B
	// It() (float64, float64, float64, float64) // present
	// X11 named colors (through RGB)
	// See also ByName(name string) fn
	Named() bool          // does current color has x11 name
	Name() string         // get name of current color or "" if haven't
	ClosestNamed() *Color // get closest named color
	// For web (through RGB)
	Websafe() string
	ShortWebsafe() string
	// For web (through RGB)
	Hex() string   // #023afe
	Short() string // #3df
	// Exotic Hex with alpha (RGBA, HSLA) (through RGBA)
	// For colors without alpha channel, it always would be 1 (alpha = 0..1)
	Hexa() string   // #RRGGBBAA for colors with alpha channel - RGBA and HSLA
	Shorta() string // #3df9 for colors with alpha channel - RGBA and HSLA
	// Currently not implemented
	Grey() *Color
	// Harmonies (through LCH)
	Tradic() (*Color, *Color)        // H +/-120°
	Split() (*Color, *Color)         // H +/-150° split complements
	Analogous() (*Color, *Color)     // H +/-30°
	Monochromatic() (*Color, *Color) // C*
	Complement() *Color              // H +180°
	// Delta (through LAB)
	DeltaC(*Color) float64
	DeltaH(*Color) float64
	DeltaE(*Color) float64
	DeltaE94(*Color, ...float64) float64
	DeltaE00(*Color, ...float64) float64
	DeltaCMC(*Color, ...float64) float64
}

type RGB struct {
	R, G, B float64
}

type RGBA struct {
	R, G, B, A float64
}

type CMY struct {
	C, M, Y float64
}

type CMYK struct {
	C, M, Y, K float64
}

type HSL struct {
	H, S, L float64
}

type HSLA struct {
	H, S, L, A float64
}

type AHSL struct {
	H, S, L float64
}

type HSB struct {
	H, S, B float64
}

type HSV struct {
	H, S, V float64
}

type LAB struct {
	L, A, B float64
}

type HunterLAB struct {
	L, A, B float64
}

type LCH struct {
	L, C, H float64
}

type LUV struct {
	L, U, V float64
}

type XYZ struct {
	X, Y, Z float64
}

type YXY struct {
	Y1, X, Y2 float64
}

// Y'CbCr
type YCbCr struct {
	Y, Cb, Cr float64 // Y', Cb, Cr
}

type YIQ struct {
	Y, I, Q float64
}

type YUV struct {
	Y, U, V float64
}

func FromHex(hex string) (*RGB, error) {
	if hex[0] == '#' { // strings.HasPrefix(hex, "#")
		hex = hex[1:]
	}
	var sr, sg, sb string
	if len(hex) == 6 {
		sr = hex[:2]
		sg = hex[2:4]
		sb = hex[4:]
	} else if len(hex) == 3 {
		sr = hex[:1]
		sg = hex[1:2]
		sb = hex[2:]
	} else {
		return nil, errors.New("Only #RRGGBB, RRGGBB, #RGB and RGB values are supported, length out of range")
	}
	ir, err := strconv.ParseInt(sr, 16, 64) // int64
	if err != nil {
		return nil, err
	}
	ig, err := strconv.ParseInt(sg, 16, 64) // int64
	if err != nil {
		return nil, err
	}
	ib, err := strconv.ParseInt(sb, 16, 64) // int64
	if err != nil {
		return nil, err
	}
	rgb := &RGB{
		float64(ir),
		float64(ig),
		float64(ib),
	}
	return rgb
}

func FromHexa(hexa string) *RGBA {
	if hex[0] == '#' { // strings.HasPrefix(hex, "#")
		hex = hex[1:]
	}
	var rgb *RGB
	var err error
	var sa string
	if len(hex) == 8 {
		rgb, err = FromHex(hex[:6])
		sa = hex[6:]
	} else if len(hex) == 4 {
		rgb, err = FromHex(hex[:3])
		sa = hex[3:]
	} else {
		return nil, errors.New("Only #RRGGBBAA, RRGGBBAA, #RGBA and RGBA values are supported, length out of range")
	}
	if err != nil {
		return nil, err
	}
	var ia int64
	ia, err = strconv.ParseInt(sa, 16, 64) // int64
	if err != nil {
		return nil, err
	}
	return rgb.AddAlpha(float64(ia))
}
