package heyday

import (
	"errors"
	"fmt" // DEBUG
)

// helpers for find inverse matrix

type Matrix3x3 [9]float64

type Matrix2x2 [4]float64

func (m *Matrix3x3) Set(a11, b11, c11, a21, b21, c21, a22, b22, c22 float64) *Matrix3x3 {
	(*m)[0], (*m)[1], (*m)[2], (*m)[3], (*m)[4], (*m)[5], (*m)[6], (*m)[7], (*m)[8] = a11, b11, c11, a21, b21, c21, a22, b22, c22
	return m
}

func (m *Matrix2x2) Set(a11, b11, a21, b21 float64) *Matrix2x2 {
	(*m)[0], (*m)[1], (*m)[2], (*m)[3] = a11, b11, a21, b21
	return m
}

func (m *Matrix3x3) Det() float64 {
	return (*m)[0]*(*m)[4]*(*m)[8] + (*m)[6]*(*m)[1]*(*m)[5] + (*m)[3]*(*m)[7]*(*m)[2] - (*m)[6]*(*m)[4]*(*m)[2] - (*m)[0]*(*m)[7]*(*m)[5] - (*m)[3]*(*m)[1]*(*m)[8]
}

func (m *Matrix2x2) Det() float64 {
	return (*m)[0]*(*m)[3] - (*m)[2]*(*m)[1]
}

func (m *Matrix3x3) Minorij(i, j int) *Matrix2x2 {
	t := new(Matrix2x2)
	switch {
	case i == 0 && j == 0:
		return t.Set((*m)[4], (*m)[5], (*m)[7], (*m)[8])
	case i == 0 && j == 1:
		return t.Set((*m)[3], (*m)[5], (*m)[6], (*m)[8])
	case i == 0 && j == 2:
		return t.Set((*m)[3], (*m)[4], (*m)[6], (*m)[7])
	case i == 1 && j == 0:
		return t.Set((*m)[1], (*m)[2], (*m)[7], (*m)[8])
	case i == 1 && j == 1:
		return t.Set((*m)[0], (*m)[2], (*m)[6], (*m)[8])
	case i == 1 && j == 2:
		return t.Set((*m)[0], (*m)[1], (*m)[6], (*m)[7])
	case i == 2 && j == 0:
		return t.Set((*m)[1], (*m)[2], (*m)[4], (*m)[5])
	case i == 2 && j == 1:
		return t.Set((*m)[0], (*m)[2], (*m)[3], (*m)[5])
	case i == 2 && j == 2:
		return t.Set((*m)[0], (*m)[1], (*m)[3], (*m)[4])
	default:
		return t
	}
}

//        | 0,0   |  0,1   |  0,2   |  1,0   |  1,1   |  1,2   |  2,0   |  2,1   |  2,2
// 0 1 2  | x x x |  x x x |  x x x |  x 1 2 |  0 x 2 |  0 1 x |  x 1 2 |  0 x 2 |  0 1 x
// 3 4 5  | x 4 5 |  3 x 5 |  3 4 x |  x x x |  x x x |  x x x |  x 4 5 |  3 x 5 |  3 4 x
// 6 7 8  | x 7 8 |  6 x 8 |  6 7 x |  x 7 8 |  6 x 8 |  6 7 x |  x x x |  x x x |  x x x

func (m *Matrix3x3) Minor() *Matrix3x3 {
	t := new(Matrix3x3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			t[i*3+j], err = m.Minorij(i, j).Det()
		}
	}
	return t
}

func (m *Matrix3x3) Dop() *Matrix3x3 {
	(*m)[1], (*m)[3], (*m)[5], (*m)[7] = -(*m)[1], -(*m)[3], -(*m)[5], -(*m)[7]
	return m
}

//  x -x  x
// -x  x -x
//  x -x  x

func (m *Matrix3x3) Trans() *Matrix3x3 {
	t := new(Matrix3x3)
	(*t)[0], (*t)[1], (*t)[2] = (*m)[0], (*m)[3], (*m)[6]
	(*t)[3], (*t)[4], (*t)[5] = (*m)[1], (*m)[4], (*m)[7]
	(*t)[6], (*t)[7], (*t)[8] = (*m)[2], (*m)[5], (*m)[8]
	return t
}

// 0 1 2    0 3 6
// 3 4 5 => 1 4 7
// 6 7 8    2 5 8

func (m *Matrix3x3) Mul(c float64) *Matrix3x3 {
	(*m)[0], (*m)[1], (*m)[2] = c*(*m)[0], c*(*m)[1], c*(*m)[2]
	(*m)[3], (*m)[4], (*m)[5] = c*(*m)[3], c*(*m)[4], c*(*m)[5]
	(*m)[6], (*m)[7], (*m)[8] = c*(*m)[6], c*(*m)[7], c*(*m)[8]
	return m
}

// multiplication: [M]*c

func (m *Matrix3x3) Div(c float64) *Matrix3x3 {
	(*m)[0], (*m)[1], (*m)[2] = (*m)[0]/c, (*m)[1]/c, (*m)[2]/c
	(*m)[3], (*m)[4], (*m)[5] = (*m)[3]/c, (*m)[4]/c, (*m)[5]/c
	(*m)[6], (*m)[7], (*m)[8] = (*m)[6]/c, (*m)[7]/c, (*m)[8]/c
	return m
}

// division: [M]/c

func (m *Matrix3x3) Inverse() (*Matrix3x3, error) {
	det := m.Det()
	if det == 0 {
		return nil, errors.New("Could't find inverse matrix, det = 0")
	}
	return m.Minor().Dop().Trans().Div(det), nil

}

// return inverse matrix of current if it's possible

// DEBUG
func (m *Matrix3x3) Show() {
	fmt.Println("Matrix 3x3")
	fmt.Println((*m)[0], (*m)[1], (*m)[2])
	fmt.Println((*m)[3], (*m)[4], (*m)[5])
	fmt.Println((*m)[6], (*m)[7], (*m)[8])
	fmt.Println()
}
