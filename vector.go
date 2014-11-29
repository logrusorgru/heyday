package heyday

type Colors interface {
	// colors
	Labs() *Labs
	XYZs() *XYZs
	HunterLabs() *HunterLabs
	Luvs() *Luvs
	LCHuvs() *LCHuvs
	LCHabs() *LCHabs
	RGBs() *RGBs
	// deltas through CIE L*a*b*
	DeltasC(*Colors) []float64
	DeltasH(*Colors) []float64
	DeltasE(*Colors) []float64
	DeltasE94(*Colors, ...float64) []float64
	DeltasE00(*Colors, ...float64) []float64
	DeltasCMC(*Colors, ...float64) []float64
}

type XYZs []XYZ

type Yxys []Yxy

type Labs []Lab

type Luvs []Luv

type LCHuvs []LCHuv

type LCHabs []LCHab

type HunterLabs []HunterLab

type RGBs []RGB
