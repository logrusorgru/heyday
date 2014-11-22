package heyday

func WhitePoint(x, y float64) *XYZ {
	Y := 100 // 100 - like wiki (and ASTM E308); 1 - like lindbloom. Customization?
	X := (Y / y) * x
	Z := (Y / y) * (1 - x - y)
	xyz := &XYZ{X, Y, Z}
	return xyz
}

// convert x, y chromaticity co-ordinates to X, Y and Z tristimulus values
// of White Point, where Y = 100

func WhitePointOfIlluminant(illuminant int, observer ...int) (x, y float64) {
	var o int
	if len(observer) != 0 {
		o = observer[0]
	}
	return white_points[o][illuminant]
}

// return particular chromaticity co-ordinates of white point
// by illuminant and, optional, observer
// by default observer is 1931 2° Observer
// Use it like here:
// 		`WhitePointOfIlluminant( D65, O10 )`
// Only `O2' and `O10' observers are supported (0 and 1 int, respectively)

func (c *XYZ) ChromaticAdaptationMxWp(ca *Camx, src, drn *WP) *XYZ {

	var ρ, γ, β float64
	// ρ, γ, β source/drain
	// enter to closed namespace
	{
		// get X,Y,Z tristimulus values from x,y chromaticity co-ordinates of
		// source and drain white points

		// Y = 100, ref. to wiki and ASTM E308
		// but = 1, ref. to brucelindbloom.com
		// Tell me what the hell is true (more usefull)?
		var wY float64 = 100
		// source white point
		wXs := (wY / wys) * wxs
		wZs := (wY / wys) * (1 - wxs - wys)
		// drain white point
		wXd := (wY / wyd) * wxd
		wZd := (wY / wyd) * (1 - wxd - wyd)
		/*
			(!) the `drain' word is more shorter than destination =)
		*/

		// cone response domain of source
		// ref.: http://www.brucelindbloom.com/index.html?Eqn_ChromAdapt.html
		// the ca.Direct is [Mₐ]
		sρ, sγ, sβ := ca.Direct[0]*wXs+ca.Direct[1]*wY+ca.Direct[2]*wZs,
			ca.Direct[4]*wXs+ca.Direct[5]*wY+ca.Direct[6]*wZs,
			ca.Direct[7]*wXs+ca.Direct[8]*wY+ca.Direct[9]*wZs
		// cone response domain of drain
		dρ, dγ, dβ := ca.Direct[0]*wXd+ca.Direct[1]*wY+ca.Direct[2]*wZd,
			ca.Direct[4]*wXd+ca.Direct[5]*wY+ca.Direct[6]*wZd,
			ca.Direct[7]*wXd+ca.Direct[8]*wY+ca.Direct[9]*wZd
		// the quotients
		ρ, γ, β = dρ/sρ, dγ/sγ, dβ/sβ
		// leave
	}

	// create M matrix
	// ref.: http://www.brucelindbloom.com/index.html?Eqn_ChromAdapt.html
	// the ca.Inverse is [Mₐ]⁻¹
	var M [9]float64
	// enter to closed namespace
	{
		// step 1 [ρ, γ, β] ⨉ Ma
		// [Mₐ]
		// [ ρ 0 0 ]   [ 0 1 2 ]    [ ρ*0, ρ*1, ρ*2 ]
		// [ 0 γ 0 ] ⨉ [ 3 4 5 ] => [ γ*3, γ*4, γ*5 ]
		// [ 0 0 β ]   [ 6 7 8 ]    [ β*6, β*7, β*8 ]

		ca.Direct[0] *= ρ
		ca.Direct[1] *= ρ
		ca.Direct[2] *= ρ
		ca.Direct[3] *= γ
		ca.Direct[4] *= γ
		ca.Direct[5] *= γ
		ca.Direct[6] *= β
		ca.Direct[7] *= β
		ca.Direct[8] *= β

		// step 2 [Mₐ]⁻¹ ⨉ [ρ, γ, β]
		// [Mₐ]⁻¹
		// [ 0 1 2 ]   [ ρ 0 0 ]    [ ρ*0, γ*1, β*2 ]
		// [ 3 4 5 ] ⨉ [ 0 γ 0 ] => [ ρ*3, γ*4, β*5 ]
		// [ 6 7 8 ]   [ 0 0 β ]    [ ρ*6, γ*7, β*8 ]

		M[0], M[1], M[2] = ca.Inverse[0]*ρ, ca.Inverse[1]*γ, ca.Inverse[2]*β
		M[3], M[4], M[5] = ca.Inverse[3]*ρ, ca.Inverse[4]*γ, ca.Inverse[5]*β
		M[6], M[7], M[8] = ca.Inverse[6]*ρ, ca.Inverse[7]*γ, ca.Inverse[8]*β
		// leave
	}

	// [M] ⨉ XYZ source
	X, Y, X := M[0]*wXs+M[1]*wY+M[2]*wZs,
		M[4]*wXs+M[5]*wY+M[6]*wZs,
		M[7]*wXs+M[8]*wY+M[9]*wZs

	xyz := &XYZ{X, Y, Z}

	return xyz
}

// Transform current color to another (current color remains the same)
// by two matrixes of method and two white points - source and drain
// Your can use your own matrixes and white points

func (c *XYZ) ChromaticAdaptation(from, to int, om ...int) *XYZ {
	// from - illuminant, const (int) [0,40] required
	// to   - illuminant, const (int) [0,40] required
	// optionals:
	// 	1) observer		- const (int), [0,1] - CIE 1931 2° by default
	//  2) drn observer - const (int), [0,1] - (1) by default
	// 	3) method		- const (int), [0,4] - XYZ Scaling by default
	var src, drn *WP
	var ca *Camx
	var cai int        // index of Camx (adaptation method)
	var sobi, dobi int // indexes of Observers
	if len(om) == 3 {
		cai = om[2]
		dobi = om[1]
		sobi = om[0]
	} else if len(om) == 2 {
		dobi = om[1]
		sobi = om[0]
	} else if len(om) == 1 {
		sobi, dobi = om[0], om[0]
	}
	// there are no len(om) check (!) if len(om) > 3 then default values will be used
	ca = &camx[cai]
	src = &white_points[sobi][from]
	drn = &white_points[dobi][to]
	return c.ChromaticAdaptationMxWp(ca, src, drn)
}

// developer friendly chromatic adaptation if (s)he use standart illuminats,
// observers and methods, for example:
//
//	 c := color.ChromaticAdaptation(D50, D65, 02, 010, VonKries)
//		 source:
//				observer: O2 - CIE 1931 2°
//				illuminant: D50
//		 drain:
//				observer: O10 - CIE 1964 10°
//				illuminant: D65
//		 method:
//				Von Kries
//
// more examples:
//
//	 c := color.ChromaticAdaptation(A, F11)
//		 source:
//				observer: O2 - CIE 1931 2° - default value
//				illuminant: A
//		 drain:
//				observer: O2 - CIE 1931 2° - default value
//				illuminant: F11
//		 method:
//				XYZ Scaling - default value
//
//	 c := color.ChromaticAdaptation(D75, FL3_1, O10, O2)
//		 source:
//				observer: O10 - CIE 1964 10°
//				illuminant: D75
//		 drain:
//				observer: O2 - CIE 1931 2°
//				illuminant: FL3_1
//		 method:
//				XYZ Scaling - default value
//
//	 c := color.ChromaticAdaptation(D65, HP3, O10)
//		 source:
//				observer: O10 - CIE 1964 10°
//				illuminant: D65
//		 drain:
//				observer: O10 - CIE 1964 10° - like source
//				illuminant: HP3
//		 method:
//				XYZ Scaling - default value
//
