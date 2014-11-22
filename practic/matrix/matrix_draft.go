package matrix

import (
//"fmt" DEBUG
)

// matrix type for math
type Matrix [][]float64

// Create column matrix from 3 values, like here
// 				[X]
// X, Y, Z =>	[Y]
// 				[Z]
func Column(x, y, z float64) *Matrix {
	t := Matrix{
		[]float64{x},
		[]float64{y},
		[]float64{z},
	}
	return &t
}

// Create identity matrix from 3 values, like here
// 				[ρ, 0, 0]
// ρ, γ, β => 	[0, γ, 0]
// 				[0, 0, β]
func Identity(ρ, γ, β float64) *Matrix {
	t := Matrix{
		[]float64{ρ, 0, 0},
		[]float64{0, γ, 0},
		[]float64{0, 0, β},
	}
	return &t
}

// Create column matrix with X, Y, Z tristimuls values by
// x, y chromatic co-ordinates of white point;
// the Y is 100
func ColumnXYZ(x, y float64) *Matrix {
	var X, Y, Z float64
	Y = 100 // 100 - like wiki (and ASTM E308); 1 - like lindbloom. Customization?
	X = (Y / y) * x
	Z = (Y / y) * (1 - x - y)
	t := Matrix{
		[]float64{X},
		[]float64{Y},
		[]float64{Z},
	}
	return &t
}

// create new matrix from empty, just append []float64 slices-rows
// t := New( []float64{1,2,3}, []foat64{4,5,6}, []float64{7,8,9} )
// 		[ 1, 2, 3 ]
//  => 	[ 4, 5, 6 ]
//		[ 7, 8, 9 ]
func (m *Matrix) New(r ...[]float64) *Matrix {
	for i := 0; i < len(r); i++ {
		*m = append(*m, r[i])
	}
	return m
}

// https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BD%D0%BE%D0%B6%D0%B5%D0%BD%D0%B8%D0%B5_%D0%BC%D0%B0%D1%82%D1%80%D0%B8%D1%86#.D0.9E.D0.BF.D1.80.D0.B5.D0.B4.D0.B5.D0.BB.D0.B5.D0.BD.D0.B8.D0.B5
// Matrix-matrix multiplication
func Mul(a, b *Matrix) *Matrix {
	var c Matrix      // result
	m := len(*a)      // height of a (rows)
	n := len((*a)[0]) // have to be == len(b)
	if len(*b) != n { // length of a (columns) and height of b (rows)
		return &c
	}
	q := len((*b)[0]) // length of b (columns)
	for i := 0; i < m; i++ {
		c = append(c, make([]float64, q))
		for j := 0; j < q; j++ {
			for r := 0; r < n; r++ {
				c[i][j] += (*a)[i][r] * (*b)[r][j]
			}
		}
	}
	return &c
}

// first matrix should be XYZ Column matrix,
// second should be [Mₐ] matrix,
// see here <http://www.brucelindbloom.com/index.html?Eqn_ChromAdapt.html>
// for details
// The result is Column matrix, it's contains a cone response domain (ρ, γ, β)
// of Column  matrix with XYZ white point
// [ρ]
// [γ]
// [β]
// This fn is nothing more than a Mul() vice versa
func ConeResponseDomain(xyz, Mₐ *Matrix) *Matrix {
	return Mul(Mₐ, m)
}

// create identity matrix by two column matrix with cone response domains
// like here
// ⒮ - source
// ⒟ - destination
// [ρ⒟]  [ρ⒮] 		[ρ⒟/ρ⒮    0       0  ]
// [γ⒟]  [γ⒮] => 	[   0   γ⒟/γ⒮     0  ]
// [β⒟], [β⒮] 		[   0      0    β⒟/β⒮]
func ConeResponseDomainsIdentity(s, d *Matrix) *Matrix {
	return Identity(
		(*d)[0][0]/(*s)[0][0],
		(*d)[0][1]/(*s)[0][1],
		(*d)[0][2]/(*s)[0][2],
	)
}

// return [M] transformaation matrix for XYZ Chromatic Adaptation
// Mₐ  - [Mₐ] 	matrix
// vMₐ - [Mₐ]⁻¹ matrix
// ws  - Column matrix with source XYZ white point tristimuls values
// wd  - Column matrix with destination XYZ white point tristimuls values
func ChromaticAdaptationMatrix(Mₐ, vMₐ, ws, wd *Matrix) *Matrix {
	scrd := ConeResponseDomain(ws, Mₐ)
	dcrd := ConeResponseDomain(wd, Mₐ)
	sd := ConeResponseDomainsIdentity(scrd, dcrd)
	ffwd := Mul(sd, Mₐ)
	back := Mul(vMₐ, ffwd)
	return back
	/*
		return Mul(
			vMₐ,
			Mul(
				ConeResponseDomainsIdentity(
					ConeResponseDomain(
						ColumnXYZ(wsx, wsy), // white point chromatic co-ords
						Mₐ,
					),
					ConeResponseDomain(
						ColumnXYZ(wdx, wdy), // white point chromatic co-ords
						Mₐ,
					),
				),
				Mₐ,
			),
		)
	*/
}

// perfom chromatic adaptation by matrix
// s - source color Column matrix
// M - adaptation matrix
func ChromaticAdaptation(M, s *Matrix) *Matrix {
	return Mul(M, s)
}

/* DEBUG */
/*
func Show( m *Matrix ) {
	fmt.Println("M:")
	for i := 0; i < len(*m); i++ {
		fmt.Println((*m)[i])

	}
	fmt.Println()
}
*/

/*

всего-то ребуется превратить 3 значения в другие три значения,
на воде

X, Y, Z - источник
Ma, Ma-1 - две матрицы, девять значений, меотд
wsx, wsy - хроматические координаты белой точки источника
wdx, wdy - назначения

на выходе
Xd, Yd, Zd - тюнненый цвет

и пиздец!

А нужно то ввсего-то ничего:
1) умножение матрицы [x]
2) преобразование хроматических координат в XYZ [x]
3) три числа в диагональную матрицу

func ChromaticAdaptation( source_x,
						  source_y,
						  source_z float64,
						  Ma, Mav [9]float64,
						  wxs, wys,
						  wxd, wyd float64 ) (x, y, z float64 ) {
	var ρ, γ, β float64
	// ρ, γ, β source/drain
	{
		var wY float64 = 100
		wXs := (wY / wys) * wxs
		wZs := (wY / wys) * (1 - wxs - wys)
		wXd := (wY / wyd) * wxd
		wZd := (wY / wyd) * (1 - wxd - wyd)
		sρ, sγ, sβ := Ma[0]*wXs + Ma[1]*wY + Ma[2]*wZs,
					  Ma[4]*wXs + Ma[5]*wY + Ma[6]*wZs,
					  Ma[7]*wXs + Ma[8]*wY + Ma[9]*wZs
		dρ, dγ, dβ := Ma[0]*wXd + Ma[1]*wY + Ma[2]*wZd,
					  Ma[4]*wXd + Ma[5]*wY + Ma[6]*wZd,
					  Ma[7]*wXd + Ma[8]*wY + Ma[9]*wZd
		ρ, γ, β = dρ/sρ, dγ/sγ, dβ/sβ
	}
	// create M matrix
	var M [9]float64
	{
		// Ma
		// [ ρ 0 0 ]   [ 0 1 2 ]    [ ρ*0, ρ*1, ρ*2 ]
		// [ 0 γ 0 ] x [ 3 4 5 ] => [ γ*3, γ*4, γ*5 ]
		// [ 0 0 β ]   [ 6 7 8 ]    [ β*6, β*7, β*8 ]

		// step 1 [ρ, γ, β] x Ma

		Ma[0] *= ρ; Ma[1] *= ρ; Ma[2] *= ρ;
		Ma[3] *= γ; Ma[4] *= γ; Ma[5] *= γ;
		Ma[6] *= β; Ma[7] *= β; Ma[8] *= β;

		// Mav
		// [ 0 1 2 ]   [ ρ 0 0 ]    [ ρ*0, γ*1, β*2 ]
		// [ 3 4 5 ] x [ 0 γ 0 ] => [ ρ*3, γ*4, β*5 ]
		// [ 6 7 8 ]   [ 0 0 β ]    [ ρ*6, γ*7, β*8 ]

		// step 2 Mav x [ρ, γ, β]

		M[0], M[1], M[2] = Ma[0]*ρ, Ma[1]*γ, Ma[2]*β
		M[3], M[4], M[5] = Ma[3]*ρ, Ma[4]*γ, Ma[5]*β
		M[6], M[7], M[8] = Ma[6]*ρ, Ma[7]*γ, Ma[8]*β
	}
	// [M] x XYZ source
	X, Y, X := M[0]*wXs + M[1]*wY + M[2]*wZs,
			   M[4]*wXs + M[5]*wY + M[6]*wZs,
			   M[7]*wXs + M[8]*wY + M[9]*wZs

	return X,Y,Z
}

*/
