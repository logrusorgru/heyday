package heyday

type CC struct { // cromaticity coordinates
	X, Y float64
}

type Color interface {
	// colors
	Lab() *Lab
	XYZ() *XYZ
	HunterLab() *HunterLab
	Luv() *Luv
	LCHuv() *LCHuv
	LCHab() *LCHab
	RGB() *RGB
	// delta through CIE L*a*b*
	DeltaC(*Color) float64
	DeltaH(*Color) float64
	DeltaE(*Color) float64
	DeltaE94(*Color, ...float64) float64
	DeltaE00(*Color, ...float64) float64
	DeltaCMC(*Color, ...float64) float64
}

type XYZ struct {
	// CIE 1931 Color Space, also known as CIE XYZ
	X, Y, Z float64
}

type Yxy struct {
	Y, Xc, Yc float64
}

type Lab struct {
	// CIE L*,a*,b*
	L, A, B float64
}

type Luv struct {
	// CIE L*,u*,v*
	L, U, V float64
}

type LCHuv struct {
	// CIE L*CH°(uv)
	L, C, H float64
}

type LCHab struct {
	// CIE L*CH°(ab)
	L, C, H float64
}

type HunterLab struct {
	// Hunter Lab
	// ref.: http://fp.optics.arizona.edu/opti588/reading/CIE_Lab_color_space.pdf
	L, A, B float64
}

type RGB struct {
	R, G, B float64
}
