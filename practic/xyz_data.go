package heyday

/*

	Standart Illuminants and Obervers

	sources:

		main:
			http://nbviewer.ipython.org/github/colour-science/colour-ipython/blob/master/notebooks/colorimetry/illuminants.ipynb
		sub:
			http://en.wikipedia.org/wiki/Standard_illuminant#White_points_of_standard_illuminants
			https://law.resource.org/pub/us/cfr/ibr/003/cie.15.2004.tables.xls


	// Now, turn off word wrap =)


Name				CIE 1931 2°										CIE 1964 10°					CCT (K)	Note
		x2						y2						x10						y10
A		0.44757					0.40745					0.45117					0.40594				2856	Incandescent / Tungsten
B		0.34842					0.35161					0.3498					0.3527				4874	{obsolete} Direct sunlight at noon
C		0.31006					0.31616					0.31039					0.31905				6774	{obsolete} Average / North sky Daylight
D50		0.34567					0.3585					0.34773					0.35952				5003	Horizon Light. ICC profile PCS
D55		0.33242					0.34743					0.33411					0.34877				5503	Mid-morning / Mid-afternoon Daylight
D60		0.32168					0.33767					0.3229915227773675		0.3391283129096501
D65		0.31271					0.32902					0.31382					0.331				6504	Noon Daylight: Television, sRGB color space
D75		0.29902					0.31485					0.29968					0.3174				7504	North sky Daylight
E		1/3						1/3						1/3						1/3					5454	Equal energy
F1		0.3131					0.33727					0.31811					0.33559				6430	Daylight Fluorescent
F2		0.37208					0.37529					0.37925					0.36733				4230	Cool White Fluorescent
F3		0.4091					0.3943					0.41761					0.38324				3450	White Fluorescent
F4		0.44018					0.40329					0.4492					0.39074				2940	Warm White Fluorescent
F5		0.31379					0.34531					0.31975					0.34246				6350	Daylight Fluorescent
F6		0.3779					0.38835					0.3866					0.37847				4150	Lite White Fluorescent
F7		0.31292					0.32933					0.31569					0.3296				6500	D65 simulator, Daylight simulator
F8		0.34588					0.35875					0.34902					0.35939				5000	D50 simulator, Sylvania F40 Design 50
F9		0.37417					0.37281					0.37829					0.37045				4150	Cool White Deluxe Fluorescent
F10		0.34609					0.35986					0.3509					0.35444				5000	Philips TL85, Ultralume 50
F11		0.38052					0.37713					0.38541					0.37123				4000	Philips TL84, Ultralume 40
F12		0.43695					0.40441					0.44256					0.39717				3000	Philips TL83, Ultralume 30
FL3.1	0.44067505367083887		0.4032982400521268		0.44983072060097606		0.3902314548749185	2932	New set of fluorescent lamps
FL3.2	0.3807750955454208		0.37335593850329946		0.3869241790592931		0.3657561399067915	3965	New set of fluorescent lamps
FL3.3	0.315282999959924		0.3438607529477053		0.3211770463590012		0.3405012294853132	6280	New set of fluorescent lamps
FL3.4	0.4429092105157883		0.40432363641263036		0.44812132825245227		0.39707718739452286	2904	New set of fluorescent lamps
FL3.5	0.37489860509162237		0.36715721144008995		0.3778142145280147		0.36662585171298506	4086	New set of fluorescent lamps
FL3.6	0.34880430112004895		0.36000066745930004		0.35197652659197876		0.3610945283215212	4894	New set of fluorescent lamps
FL3.7	0.4384254677714265		0.4045310539294221		0.44430926608426524		0.3967914695376695	2979	New set of fluorescent lamps
FL3.8	0.38197740008964315		0.3831793433643612		0.3875889785058527		0.37630564994236454	4006	New set of fluorescent lamps
FL3.9	0.3498569555988198		0.3590880129881551		0.3546890235509348		0.3534450969379135	4853	New set of fluorescent lamps
FL3.10	0.345504278345313		0.3559512432054011		0.3493448005253757		0.35498443751900227	5000	New set of fluorescent lamps
FL3.11	0.3245099692216317		0.343369788104288		0.32926800293079544		0.33886544553697306	5854	New set of fluorescent lamps
FL3.12	0.4376746480519804		0.40366659242713887		0.4422521047116721		0.40122058660685506	2984	New set of fluorescent lamps
FL3.13	0.38305207018919446		0.37244321884144194		0.3862752927999335		0.3742832325728946	3896	New set of fluorescent lamps
FL3.14	0.3447222271025105		0.3609370209728111		0.34725510277403043		0.3668082930976077	5045	New set of fluorescent lamps
FL3.15	0.3126706172769415		0.3287545660713704		0.31461426562662165		0.3333777810577307	6509	New set of fluorescent lamps
HP1		0.5330008227923825		0.4149532376869364		0.543334641014127		0.40528934193733374	1959	High pressure discharge lamps
HP2		0.4777920749861397		0.4158404588084794		0.48264737821938364		0.4108157059894971	2506	High pressure discharge lamps
HP3		0.4302314003272559		0.4075170114048384		0.4355600804884421		0.39880115305546204	3144	High pressure discharge lamps
HP4		0.38117189471910373		0.3797265745279147		0.3851937462655714		0.3682756593039444	4002	High pressure discharge lamps
HP5		0.37758320909195225		0.3713479728022603		0.38031641910835223		0.3666171209128104	4039	High pressure discharge lamps

*/

// White Points for Standart Illuminants

const (
	A int = iota
	B
	C
	D50
	D55
	D60
	D65
	D75
	E
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	FL3_1
	FL3_2
	FL3_3
	FL3_4
	FL3_5
	FL3_6
	FL3_7
	FL3_8
	FL3_9
	FL3_10
	FL3_11
	FL3_12
	FL3_13
	FL3_14
	FL3_15
	HP1
	HP2
	HP3
	HP4
	HP5
)

// D65 is default

// Observers

const (
	O2 int = iota
	O10
)

// O2 is CIE 1931 2° Observer | default
// O10 is CIE 1964 10° Observer

// White Point type
// You may also use your own white points for chromatic adaptation

type WP struct {
	X, Y float64
}

// contains chromaticity co-ordinates of standart white points

var white_points = [2][41]WP{
	// ??? fuck it ??? => instead of a [2][41]WP use slices, becaus of slice is reference type but array is not

	// O2 is CIE 1931 2° Observer | default
	[41]WP{
		WP{0.44757, 0.40745},                         /*A*/
		WP{0.34842, 0.35161},                         /*B*/
		WP{0.31006, 0.31616},                         /*C*/
		WP{0.34567, 0.3585},                          /*D50*/
		WP{0.33242, 0.34743},                         /*D55*/
		WP{0.32168, 0.33767},                         /*D60*/
		WP{0.31271, 0.32902},                         /*D65*/
		WP{0.29902, 0.31485},                         /*D75*/
		WP{1 / 3, 1 / 3},                             /*E*/
		WP{0.3131, 0.33727},                          /*F1*/
		WP{0.37208, 0.37529},                         /*F2*/
		WP{0.4091, 0.3943},                           /*F3*/
		WP{0.44018, 0.40329},                         /*F4*/
		WP{0.31379, 0.34531},                         /*F5*/
		WP{0.3779, 0.38835},                          /*F6*/
		WP{0.31292, 0.32933},                         /*F7*/
		WP{0.34588, 0.35875},                         /*F8*/
		WP{0.37417, 0.37281},                         /*F9*/
		WP{0.34609, 0.35986},                         /*F10*/
		WP{0.38052, 0.37713},                         /*F11*/
		WP{0.43695, 0.40441},                         /*F12*/
		WP{0.44067505367083887, 0.4032982400521268},  /*FL3.1*/
		WP{0.3807750955454208, 0.37335593850329946},  /*FL3.2*/
		WP{0.315282999959924, 0.3438607529477053},    /*FL3.3*/
		WP{0.4429092105157883, 0.40432363641263036},  /*FL3.4*/
		WP{0.37489860509162237, 0.36715721144008995}, /*FL3.5*/
		WP{0.34880430112004895, 0.36000066745930004}, /*FL3.6*/
		WP{0.4384254677714265, 0.4045310539294221},   /*FL3.7*/
		WP{0.38197740008964315, 0.3831793433643612},  /*FL3.8*/
		WP{0.3498569555988198, 0.3590880129881551},   /*FL3.9*/
		WP{0.345504278345313, 0.3559512432054011},    /*FL3.10*/
		WP{0.3245099692216317, 0.343369788104288},    /*FL3.11*/
		WP{0.4376746480519804, 0.40366659242713887},  /*FL3.12*/
		WP{0.38305207018919446, 0.37244321884144194}, /*FL3.13*/
		WP{0.3447222271025105, 0.3609370209728111},   /*FL3.14*/
		WP{0.3126706172769415, 0.3287545660713704},   /*FL3.15*/
		WP{0.5330008227923825, 0.4149532376869364},   /*HP1*/
		WP{0.4777920749861397, 0.4158404588084794},   /*HP2*/
		WP{0.4302314003272559, 0.4075170114048384},   /*HP3*/
		WP{0.38117189471910373, 0.3797265745279147},  /*HP4*/
		WP{0.37758320909195225, 0.3713479728022603},  /*HP5*/
	},
	[41]WP{
		// O10 is CIE 1964 10° Observer
		WP{0.45117, 0.40594},                         /*A*/
		WP{0.3498, 0.3527},                           /*B*/
		WP{0.31039, 0.31905},                         /*C*/
		WP{0.34773, 0.35952},                         /*D50*/
		WP{0.33411, 0.34877},                         /*D55*/
		WP{0.3229915227773675, 0.3391283129096501},   /*D60*/
		WP{0.31382, 0.331},                           /*D65*/
		WP{0.29968, 0.3174},                          /*D75*/
		WP{1 / 3, 1 / 3},                             /*E*/
		WP{0.31811, 0.33559},                         /*F1*/
		WP{0.37925, 0.36733},                         /*F2*/
		WP{0.41761, 0.38324},                         /*F3*/
		WP{0.4492, 0.39074},                          /*F4*/
		WP{0.31975, 0.34246},                         /*F5*/
		WP{0.3866, 0.37847},                          /*F6*/
		WP{0.31569, 0.3296},                          /*F7*/
		WP{0.34902, 0.35939},                         /*F8*/
		WP{0.37829, 0.37045},                         /*F9*/
		WP{0.3509, 0.35444},                          /*F10*/
		WP{0.38541, 0.37123},                         /*F11*/
		WP{0.44256, 0.39717},                         /*F12*/
		WP{0.44983072060097606, 0.3902314548749185},  /*FL3.1*/
		WP{0.3869241790592931, 0.3657561399067915},   /*FL3.2*/
		WP{0.3211770463590012, 0.3405012294853132},   /*FL3.3*/
		WP{0.44812132825245227, 0.39707718739452286}, /*FL3.4*/
		WP{0.3778142145280147, 0.36662585171298506},  /*FL3.5*/
		WP{0.35197652659197876, 0.3610945283215212},  /*FL3.6*/
		WP{0.44430926608426524, 0.3967914695376695},  /*FL3.7*/
		WP{0.3875889785058527, 0.37630564994236454},  /*FL3.8*/
		WP{0.3546890235509348, 0.3534450969379135},   /*FL3.9*/
		WP{0.3493448005253757, 0.35498443751900227},  /*FL3.10*/
		WP{0.32926800293079544, 0.33886544553697306}, /*FL3.11*/
		WP{0.4422521047116721, 0.40122058660685506},  /*FL3.12*/
		WP{0.3862752927999335, 0.3742832325728946},   /*FL3.13*/
		WP{0.34725510277403043, 0.3668082930976077},  /*FL3.14*/
		WP{0.31461426562662165, 0.3333777810577307},  /*FL3.15*/
		WP{0.543334641014127, 0.40528934193733374},   /*HP1*/
		WP{0.48264737821938364, 0.4108157059894971},  /*HP2*/
		WP{0.4355600804884421, 0.39880115305546204},  /*HP3*/
		WP{0.3851937462655714, 0.3682756593039444},   /*HP4*/
		WP{0.38031641910835223, 0.3666171209128104},  /*HP5*/
	},
}

/*

	Methods of Chromatic Adaptation

	sources:
		http://www.brucelindbloom.com/index.html?Eqn_ChromAdapt.html
		http://hrcak.srce.hr/file/95370
		http://play.golang.org/p/rRbBdVgT0j
		http://play.golang.org/p/bVgaCMfi52

			CAT			|				[Ma]			  |				  [Mₐ]⁻¹
	--------------------+---------------------------------+------------------------------------------------------------------------
						|	[  1.0000  0.0000  0.0000 ]	  |	  [  1.0000  0.0000  0.0000 ]
		XYZ Scaling		|	[  0.0000  1.0000  0.0000 ]	  |	  [  0.0000  1.0000  0.0000 ]
						|	[  0.0000  0.0000  1.0000 ]	  |	  [  0.0000  0.0000  1.0000 ]
	--------------------+---------------------------------+------------------------------------------------------------------------
						|	[  0.8951  0.2664 -0.1614 ]	  |	  [  0.986992905466712100 -0.147054256420990100  0.159962651663731220 ]
		  Bradford		|	[ -0.7502  1.7135  0.0367 ]	  |	  [  0.432305269723394500  0.518360271536777600  0.049291228212855594 ]
						|	[  0.0389  0.0685  1.0296 ]	  |	  [ -0.008528664575177328  0.040042821654084860  0.968486695787550000 ]
	--------------------+---------------------------------+------------------------------------------------------------------------
						|	[  0.4002  0.7076 -0.0808 ]	  |	  [  1.859936387455839700 -1.129381618580091600  0.219897409596193280 ]
		  Von Kries		|	[ -0.2263  1.1653  0.0457 ]	  |	  [  0.361191436241767600  0.638812463285042200 -0.000006370596838651 ]
						|	[  0.0000  0.0000  0.9182 ]	  |	  [ -0.000000000000000000 -0.000000000000000000  1.089063623096861300 ]
	--------------------+---------------------------------+------------------------------------------------------------------------
						|	[  0.8951 -0.7502  0.0389 ]	  |	  [  0.977582533226663700  0.429405537812052900 -0.065503340989357840 ]
		  CMCCAT97		|	[  0.2664  1.7135  0.0685 ]	  |	  [ -0.158337864212157560  0.514883316879461800 -0.028273275338374315 ]
						|	[ -0.1614  0.0367  1.0296 ]	  |	  [  0.158889685780273570  0.048960602246881375  0.961990471998092400 ]
	--------------------+---------------------------------+------------------------------------------------------------------------
						|	[  0.7982  0.3389 -0.1371 ]	  |	  [  1.076450048678639700 -0.237662388092563040  0.1612123394139235000 ]
		  CMCCAT200		|	[ -0.5918  1.5512  0.0406 ]	  |	  [  0.410964325479775900  0.554341804147103200  0.0346938703731209900 ]
						|	[  0.0008  0.0239  0.9753 ]	  |	  [ -0.010953765423879379 -0.013389356309486022  1.0243431217333656000 ]
	------------------------------------------------------+------------------------------------------------------------------------


*/

// Methods of Chromatic Adaptation

const (
	XYZScaling int = iota
	Bradford
	VonKries
	CMCCAT97
	CMCCAT200
)

// Chromatic Adaptetion Marixes
// Direct and Inverse
// Inverse matrixes are precomputed for perfomance gain
// There are four methods (exclusive XYZ Scale), but you can use your own matrixes for chromatic adaptation

type Camx struct {
	Direct,
	Inverse [9]float64
}

// cromatic adaptation matrixes

var camx = [5]Camx{
	// XYZ Scaling matrixes
	Camx{
		[9]float64{ // direct
			1, 0, 0,
			0, 1, 0,
			0, 0, 1,
		},
		[9]float64{ // inverse
			1, 0, 0,
			0, 1, 0,
			0, 0, 1,
		},
	},
	// Bradford
	Camx{
		[9]float64{
			0.8951, 0.2664, -0.1614,
			-0.7502, 1.7135, 0.0367,
			0.0389, 0.0685, 1.0296,
		},
		[9]float64{
			0.986992905466712100, -0.147054256420990100, 0.159962651663731220,
			0.432305269723394500, 0.518360271536777600, 0.049291228212855594,
			0.008528664575177328, 0.040042821654084860, 0.968486695787550000,
		},
	},
	// Von Kries
	Camx{
		[9]float64{
			0.4002, 0.7076, -0.0808,
			-0.2263, 1.1653, 0.0457,
			0.0000, 0.0000, 0.9182,
		},
		[9]float64{
			1.859936387455839700, -1.129381618580091600, 0.219897409596193280,
			0.361191436241767600, 0.638812463285042200, -0.000006370596838651,
			-0.000000000000000000, -0.000000000000000000, 1.089063623096861300,
		},
	},
	// CMCCAT97
	Camx{
		[9]float64{
			0.8951, -0.7502, 0.0389,
			0.2664, 1.7135, 0.0685,
			-0.1614, 0.0367, 1.0296,
		},
		[9]float64{
			0.977582533226663700, 0.429405537812052900, -0.065503340989357840,
			-0.158337864212157560, 0.514883316879461800, -0.028273275338374315,
			0.158889685780273570, 0.048960602246881375, 0.961990471998092400,
		},
	},
	// CMCCAT200
	Camx{
		[9]float64{
			0.7982, 0.3389, -0.1371,
			-0.5918, 1.5512, 0.0406,
			0.0008, 0.0239, 0.9753,
		},
		[9]float64{
			1.076450048678639700, -0.237662388092563040, 0.1612123394139235000,
			0.410964325479775900, 0.554341804147103200, 0.0346938703731209900,
			-0.010953765423879379, -0.013389356309486022, 1.0243431217333656000,
		},
	},
}
