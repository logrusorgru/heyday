package heyday

/*
	ref.:
		main: http://www.brucelindbloom.com/index.html?WorkingSpaceInfo.html#WorkingSpaceSpecifications
		sub:  http://en.wikipedia.org/wiki/RGB_color_space#Specifications


Name			Gamma Reference		  Red Primary				 Green Primary				 Blue Primary
					    White
									x		y		Y			x		y		Y			x		y		Y
Adobe RGB(1998)	 2.2	D65		0.6400	0.3300	0.297361	0.2100	0.7100	0.627355	0.1500	0.0600	0.075285
Apple RGB		 1.8	D65		0.6250	0.3400	0.244634	0.2800	0.5950	0.672034	0.1550	0.0700	0.083332
Best RGB		 2.2	D50		0.7347	0.2653	0.228457	0.2150	0.7750	0.737352	0.1300	0.0350	0.034191
Beta RGB		 2.2	D50		0.6888	0.3112	0.303273	0.1986	0.7551	0.663786	0.1265	0.0352	0.032941
Bruce RGB		 2.2	D65		0.6400	0.3300	0.240995	0.2800	0.6500	0.683554	0.1500	0.0600	0.075452
CIE RGB			 2.2	E		0.7350	0.2650	0.176204	0.2740	0.7170	0.812985	0.1670	0.0090	0.010811
ColorMatch RGB	 1.8	D50		0.6300	0.3400	0.274884	0.2950	0.6050	0.658132	0.1500	0.0750	0.066985
Don RGB 4		 2.2	D50		0.6960	0.3000	0.278350	0.2150	0.7650	0.687970	0.1300	0.0350	0.033680
ECI RGB v2		 L* 	D50		0.6700	0.3300	0.320250	0.2100	0.7100	0.602071	0.1400	0.0800	0.077679
Ekta Space PS5	 2.2	D50		0.6950	0.3050	0.260629	0.2600	0.7000	0.734946	0.1100	0.0050	0.004425
NTSC RGB		 2.2	C		0.6700	0.3300	0.298839	0.2100	0.7100	0.586811	0.1400	0.0800	0.114350
PAL/SECAM RGB	 2.2	D65		0.6400	0.3300	0.222021	0.2900	0.6000	0.706645	0.1500	0.0600	0.071334
ProPhoto RGB	 1.8	D50		0.7347	0.2653	0.288040	0.1596	0.8404	0.711874	0.0366	0.0001	0.000086
SMPTE-C RGB		 2.2	D65		0.6300	0.3400	0.212395	0.3100	0.5950	0.701049	0.1550	0.0700	0.086556
sRGB			≈2.2	D65		0.6400	0.3300	0.212656	0.3000	0.6000	0.715158	0.1500	0.0600	0.072186
Wide Gamut RGB	 2.2	D50		0.7350	0.2650	0.258187	0.1150	0.8260	0.724938	0.1570	0.0180	0.016875

*/

// RGB Color Spaces

const (
	Adobe_RGB_1998 int = iota
	Apple_RGB
	Best_RGB
	Beta_RGB
	Bruce_RGB
	CIE_RGB
	ColorMatch_RGB
	Don_RGB_4
	ECI_RGB_v2
	Ekta_Space_PS5
	NTSC_RGB
	PAL_SECAM_RGB
	ProPhoto_RGB
	SMPTE_C_RGB
	SRGB // sRGB
	Wide_Gamut_RGB
)

type Senary struct {
	Xr, Yr,
	Xg, Yg,
	Xb, Yb float64
}

var rgb_white_points = [16]Senary{
	Senary{0.6400, 0.3300, 0.2100, 0.7100, 0.1500, 0.0600},
	Senary{0.6250, 0.3400, 0.2800, 0.5950, 0.1550, 0.0700},
	Senary{0.7347, 0.2653, 0.2150, 0.7750, 0.1300, 0.0350},
	Senary{0.6888, 0.3112, 0.1986, 0.7551, 0.1265, 0.0352},
	Senary{0.6400, 0.3300, 0.2800, 0.6500, 0.1500, 0.0600},
	Senary{0.7350, 0.2650, 0.2740, 0.7170, 0.1670, 0.0090},
	Senary{0.6300, 0.3400, 0.2950, 0.6050, 0.1500, 0.0750},
	Senary{0.6960, 0.3000, 0.2150, 0.7650, 0.1300, 0.0350},
	Senary{0.6700, 0.3300, 0.2100, 0.7100, 0.1400, 0.0800},
	Senary{0.6950, 0.3050, 0.2600, 0.7000, 0.1100, 0.0050},
	Senary{0.6700, 0.3300, 0.2100, 0.7100, 0.1400, 0.0800},
	Senary{0.6400, 0.3300, 0.2900, 0.6000, 0.1500, 0.0600},
	Senary{0.7347, 0.2653, 0.1596, 0.8404, 0.0366, 0.0001},
	Senary{0.6300, 0.3400, 0.3100, 0.5950, 0.1550, 0.0700},
	Senary{0.6400, 0.3300, 0.3000, 0.6000, 0.1500, 0.0600},
	Senary{0.7350, 0.2650, 0.1150, 0.8260, 0.1570, 0.0180},
}

func RgbWhitePointOf(color_space int) *Senary {
	return &rgb_white_points[color_space]
}

// return particular rgb white point

func RgbDirectMatrix(cs *Senary, wp *WP) (*Matrix3x3, error) { // cs = color_space
	wt := TristimulusWhite(wp.X, wp.Y)
	Y := float64(1) // Y of white (1 or 100 or 255 or table cell ???)
	Xr := cs.Xr / cs.Yr
	Zr := (1 - cs.Xr - cs.Yr) / cs.Yr
	Xg := cs.Xg / cs.Yg
	Zg := (1 - cs.Xg - cs.Yg) / cs.Yg
	Xb := cs.Xb / cs.Yb
	Zb := (1 - cs.Xb - cs.Yb) / cs.Yb
	mx, err := (&Matrix3x3{}).Set(
		Xr, Xg, Xb, Y, Y, Y, Zr, Zg, Zb,
	).Inverse()
	if err != nil {
		return nil, err
	}
	Sr, Sg, Sb := mx.RightColumn(wt.X, wt.Y, wt.Z)
	return &Matrix3x3{
		Sr * Xr, Sg * Xg, Sb * Xb,
		Sr * Y, Sg * Y, Sb * Y, // if Y 100% == 1, multiplication and Y --> nah...
		Sr * Zr, Sg * Zg, Sb * Zb,
	}
}

// ref.: http://www.brucelindbloom.com/index.html?Eqn_RGB_XYZ_Matrix.html
// [X]      [R]
// [Y] = [M][G]
// [Z]      [B]
// return matrix [M]
// argumants: rgb white point (type *Senary) and XYZ white point (type *WP)

func RgbInverseMatrix(cs *Senary, wp *WP) (*Matrix3x3, error) { // cs = color_space
	return RgbDirectMatrix(cs, wp).Inverse()
}

// ref.: http://www.brucelindbloom.com/index.html?Eqn_RGB_XYZ_Matrix.html
// [X]      [R]         [R]        [X]
// [Y] = [M][G]   and   [G] = [M]⁻¹[Y]
// [Z]      [B]         [B]        [Z]
// return matrix [M]⁻¹
// argumants: rgb white point (type *Senary) and XYZ white point (type *WP)
