package heyday

import (
	"fmt" // DEBUG
	. "github.com/logrusorgru/heyday"
	"math"
)

func (c LCHabs) Labs() Labs {
	for i := 0; i < len(c); i++ {
		a := math.Cos(degree_2_radian(c[i].H)) * c[i].C
		b := math.Sin(degree_2_radian(c[i].H)) * c[i].C
		// an L is the same
		c[i].C = a // A
		c[i].H = b // B
	}
	// TRANSFORM LCHuvs to Luvs
	cc := *(*Labs)(unsafe.Pointer(&c))
	return cc
}

// transform to CIE L*,a*,b*'s from current color points

func (c LCHabs) LCHabs() LCHabs {
	return c
}

// return self

// DEBUG
func (c LCHabs) Show() {
	fmt.Println("CIE L*CH°ab Array, Length %d", len(c))
	for i := 0; i < len(c); i++ {
		fmt.Printf("L*: %.48f, C : %.48f, H°: %.48f\n", c[i].L, c[i].C, c[i].H)
	}
}
