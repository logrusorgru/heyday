package heyday

import (
	"fmt" // DEBUG
	. "github.com/logrusorgru/heyday"
	"math"
)

func (c LCHuvs) Luvs() Luvs {
	for i := 0; i < len(c); i++ {
		u := math.Cos(degree_2_radian(c[i].H)) * c[i].C
		v := math.Sin(degree_2_radian(c[i].H)) * c[i].C
		// an L is the same
		c[i].C = u // U
		c[i].H = v // V
	}
	// TRANSFORM LCHuvs to Luvs
	cc := *(*Luvs)(unsafe.Pointer(&c))
	return cc
}

// create CIE L*,u*,v* from current color point

func (c LCHuvs) LCHuvs() LCHuvs {
	return c
}

// return self

// DEBUG
func (c *LCHuv) Show() {
	fmt.Println("CIE L*CH°uv Array, Length %d", len(c))
	for i := 0; i < len(c); i++ {
		fmt.Printf("L*: %.48f, C : %.48f, H°: %.48f\n", c[i].L, c[i].C, c[i].H)
	}
}
