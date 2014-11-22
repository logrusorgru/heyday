/*
package provide colors stuff; supported colors: CIE 1931 XYZ, CIE 1964 U*V*W*,
CIE 1960 UCS, CIE 1976 L*,u*,v*, sRGB (many), Y'CbCr, YIQ, YDbDr, Y'UV,
LMS, H°SL(a), H°SV(/B), YCgCo, OSA-UCS, AH°SL, CIE L*,a*,b*, Hunter L,a,b, xyY, LCH°
Beside, x11 named colors, CMYK, CMY, web RGB(a) and HSL(a).
Calculate delta (C,E,E94,E00,CMC), chromatic adaptation (for XYZ) and more...
*/
package heyday

type XYZ struct {
	X, Y, Z float64
}
