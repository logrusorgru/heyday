package heyday

/*
package provide colors stuff; supported colors: CIE 1931 XYZ, CIE 1964 U*V*W*,
CIE 1960 UCS, CIE 1976 L*,u*,v*, sRGB (many), Y'CbCr, YIQ, YDbDr, Y'UV,
LMS, H°SL(a), H°SV(/B), YCgCo, OSA-UCS, AH°SL, CIE L*,a*,b*, Hunter L,a,b, xyY, LCH°
Beside, x11 named colors, CMYK, CMY, web RGB(a) and HSL(a).
Calculate delta (C,E,E94,E00,CMC), chromatic adaptation (for XYZ) and more...
*/

type XYZ struct {
	// CIE 1931 Color Space, also known as CIE XYZ
	X, Y, Z float64
}

type Lab struct {
	// CIE L*,a*,b*
	L, A, B float64
}

type Luv struct {
	// CIE L*,u*,v*
	L, U, V float64
}

type LCH struct {
	// CIE L*CH°
	L, C, H float64
}
