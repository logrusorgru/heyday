package heyday

import (
	. "github.com/logrusorgru/heyday"
)

func (c XYZs) ChromaticAdaptationBy(ca Camx, src, drn *WP) XYZs {
	// ca is a copy of struct

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
		wXs := (wY / src.Y) * src.X
		wZs := (wY / src.Y) * (1 - src.X - src.Y)
		// drain white point
		wXd := (wY / drn.Y) * drn.X
		wZd := (wY / drn.Y) * (1 - drn.X - drn.Y)
		/*
			(!) the `drain' word is more shorter than destination =)
		*/

		// cone response domain of source
		// ref.: http://www.brucelindbloom.com/index.html?Eqn_ChromAdapt.html
		// the ca.Direct is [Mₐ]
		sρ, sγ, sβ := ca.Direct[0]*wXs+ca.Direct[1]*wY+ca.Direct[2]*wZs,
			ca.Direct[3]*wXs+ca.Direct[4]*wY+ca.Direct[5]*wZs,
			ca.Direct[6]*wXs+ca.Direct[7]*wY+ca.Direct[8]*wZs
		// cone response domain of drain
		dρ, dγ, dβ := ca.Direct[0]*wXd+ca.Direct[1]*wY+ca.Direct[2]*wZd,
			ca.Direct[3]*wXd+ca.Direct[4]*wY+ca.Direct[5]*wZd,
			ca.Direct[6]*wXd+ca.Direct[7]*wY+ca.Direct[8]*wZd
		// the quotients
		ρ, γ, β = dρ/sρ, dγ/sγ, dβ/sβ
		// leave
	}

	// create M matrix
	// ref.: http://www.brucelindbloom.com/index.html?Eqn_ChromAdapt.html
	// the ca.Inverse is [Mₐ]⁻¹
	var M Matrix3x3
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

	// main loop

	for i := 0; i < len(c); i++ {
		// [M] ⨉ XYZ source
		x, y, z := M[0]*c[i].X+M[1]*c[i].Y+M[2]*c[i].Z,
			M[3]*c[i].X+M[4]*c[i].Y+M[5]*c[i].Z,
			M[6]*c[i].X+M[7]*c[i].Y+M[8]*c[i].Z

		c[i].X = x
		c[i].Y = y
		c[i].Z = z
	}

	return c
}

// Transform current color to another (current color remains the same)
// by two matrices of method and two white points - source and drain
// You can use your own matrices and white points

func (c XYZs) ChromaticAdaptation(from, to int, om ...int) XYZs {
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
	// there are no check of len(om) (!) if len(om) > 3 then default values will be used
	ca = &camx[cai]
	src = &white_points[sobi][from]
	drn = &white_points[dobi][to]
	return c.ChromaticAdaptationBy(*ca, src, drn)
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
