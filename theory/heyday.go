/*
package provide different color models: CIE L*a*b*, RGB(many), H°SL, H°SV(B), CMYK, CMY, XYZ, Yxy,
YIQ, Y'CbCr, AH°SL, Hunter L,a,b, LCH°, Luv, Yuv, (with alpha channel: RGBA, HSLA) X11 named colors
and stuff to work with its, including convert from-to (inclusive convert to hex, css, x11 color name,
websafe color) calculate delta (C,E,E94,E00,CMC) and more...
*/
package heyday

/*

	http://en.wikipedia.org/wiki/Standard_illuminant
	http://en.wikipedia.org/wiki/Standard_illuminant#White_points_of_standard_illuminants

	White points of standard illuminants

Name		CIE 1931 2°				CIE 1964 10°		CCT (K)		Note
		x2			y2			x10			y10
A		0.44757		0.40745		0.45117		0.40594		2856		Incandescent / Tungsten
B		0.34842		0.35161		0.34980		0.35270		4874		{obsolete} Direct sunlight at noon
C		0.31006		0.31616		0.31039		0.31905		6774		{obsolete} Average / North sky Daylight
D50		0.34567		0.35850		0.34773		0.35952		5003		Horizon Light. ICC profile PCS
D55		0.33242		0.34743		0.33411		0.34877		5503		Mid-morning / Mid-afternoon Daylight
D65		0.31271		0.32902		0.31382		0.33100		6504		Noon Daylight: Television, sRGB color space
D75		0.29902		0.31485		0.29968		0.31740		7504		North sky Daylight
E		1/3			1/3			1/3			1/3			5454		Equal energy
F1		0.31310		0.33727		0.31811		0.33559		6430		Daylight Fluorescent
F2		0.37208		0.37529		0.37925		0.36733		4230		Cool White Fluorescent
F3		0.40910		0.39430		0.41761		0.38324		3450		White Fluorescent
F4		0.44018		0.40329		0.44920		0.39074		2940		Warm White Fluorescent
F5		0.31379		0.34531		0.31975		0.34246		6350		Daylight Fluorescent
F6		0.37790		0.38835		0.38660		0.37847		4150		Lite White Fluorescent
F7		0.31292		0.32933		0.31569		0.32960		6500		D65 simulator, Daylight simulator
F8		0.34588		0.35875		0.34902		0.35939		5000		D50 simulator, Sylvania F40 Design 50
F9		0.37417		0.37281		0.37829		0.37045		4150		Cool White Deluxe Fluorescent
F10		0.34609		0.35986		0.35090		0.35444		5000		Philips TL85, Ultralume 50
F11		0.38052		0.37713		0.38541		0.37123		4000		Philips TL84, Ultralume 40
F12		0.43695		0.40441		0.44256		0.39717		3000		Philips TL83, Ultralume 30

*/

/*

	http://en.wikipedia.org/wiki/CIE_1931_color_space#Tristimulus_values
	http://en.wikipedia.org/wiki/CIE_1931_color_space#CIE_xy_chromaticity_diagram_and_the_CIE_xyY_color_space

	Chromaticity by tristimulus values

	         X
	x = -----------
	     X + Y + Z

	         Y
	y = -----------
	     X + Y + Z

	         Z
	z = -----------
	     X + Y + Z

	as    x + y + z = 1

	=>    z = 1 - x - y

*/

/*

	Tristimuls values by chromatic, where Y = 100 (white point)

          Y
	X = -----

*/

/*

	http://www.brucelindbloom.com/index.html?Eqn_ChromAdapt.html

	Chromatic Adaptation

 Method					[Mₐ]									[Mₐ]⁻¹

           	 1.0000000  0.0000000  0.0000000		 1.0000000  0.0000000  0.0000000
XYZ Scaling	 0.0000000  1.0000000  0.0000000		 0.0000000  1.0000000  0.0000000
           	 0.0000000  0.0000000  1.0000000		 0.0000000  0.0000000  1.0000000

        	 0.8951000  0.2664000 -0.1614000		 0.9869929 -0.1470543  0.1599627
Bradford	-0.7502000  1.7135000  0.0367000		 0.4323053  0.5183603  0.0492912
        	 0.0389000 -0.0685000  1.0296000		-0.0085287  0.0400428  0.9684867

         	 0.4002400  0.7076000 -0.0808100		 1.8599364 -1.1293816  0.2198974
Von Kries	-0.2263000  1.1653200  0.0457000		 0.3611914  0.6388125 -0.0000064
         	 0.0000000  0.0000000  0.9182200		 0.0000000  0.0000000  1.0890636

*/

/*
	http://www.brucelindbloom.com/index.html?Eqn_ChromAdapt.html

	Chromatic Adaptation

	⒮ - source
	⒟ - destination
	ⓦ - reference white

	[X⒟]      [X⒮]
	[Y⒟] = [M][Y⒮]
	[Z⒟]      [Z⒮]

	[ρ⒮]       [Xⓦ⒮]
	[γ⒮] = [Mₐ][Yⓦ⒮]
	[β⒮]       [Zⓦ⒮]

	[ρ⒟]       [Xⓦ⒟]
	[γ⒟] = [Mₐ][Yⓦ⒟]
	[β⒟]       [Zⓦ⒟]

                [ρ⒟/ρ⒮    0       0  ]
	[M] = [Mₐ]⁻¹[   0   γ⒟/γ⒮     0  ][Mₐ]
	            [   0      0    β⒟/β⒮]

	[M] for illuminant (A to B for example with 2° observer)

	func XYZChromaticAdaptationAtoB2°( x⒮, y⒮, z⒮ float64 ) (x⒟, y⒟, z⒟ float64) {

		//  [Mₐ] Bradford						 [Mₐ]⁻¹ Bradford
		//  0.8951000  0.2664000 -0.1614000		 0.9869929 -0.1470543  0.1599627
		// -0.7502000  1.7135000  0.0367000		 0.4323053  0.5183603  0.0492912
		//  0.0389000 -0.0685000  1.0296000		-0.0085287  0.0400428  0.9684867

		//Whites
		// this is chromatic co-ords.
		xⓦ⒜,yⓦ⒜ := 0.44757, 0.40745 // A
		xⓦ⒝,yⓦ⒝ := 0.34842, 0.35161 // B

		// calculate tristimulas values of white points
		/ *
			     Y
			X = ---x
			     y

				 Y
			Z = ---(1 - x - y)
				 y

			where: Y = 100

		* /

		Xⓦ⒜ := 100*xⓦ⒜/yⓦ⒜
		Xⓦ⒝ := 100*xⓦ⒝/yⓦ⒝

		Zⓦ⒜ = 100*(1-xⓦ⒜-yⓦ⒜)/yⓦ⒜
		Zⓦ⒝ = 100*(1-xⓦ⒝-yⓦ⒝)/yⓦ⒝

		// calc. cone response domain (ρ, γ, β) for source and dest.

		ρ⒮ := 0.8951000*Xⓦ⒜ + 0.2664000*100 - 0.1614000*Zⓦ⒜
		γ⒮ := -0.7502000*Xⓦ⒜ + 1.7135000*100 + 0.0367000*Zⓦ⒜
		β⒮ := 0.0389000*Xⓦ⒜ - 0.0685000*100 + 1.0296000*Zⓦ⒜

		ρ⒟ := 0.8951000*Xⓦ⒝ + 0.2664000*100 - 0.1614000*Zⓦ⒝
		γ⒟ := -0.7502000*Xⓦ⒝ + 1.7135000*100 + 0.0367000*Zⓦ⒝
		β⒟ := 0.0389000*Xⓦ⒝ - 0.0685000*100 + 1.0296000*Zⓦ⒝

		// calc matrix [M]

		var M [3]float64

				[matrix]       [Mₐ]		   [Mₐ]⁻¹

		M[0] = ( ρ⒮/ρ⒟ ) * 0.8951000 * 0.9869929
		M[1] = ( γ⒮/γ⒟ ) * 1.7135000 * 0.5183603
		M[2] = ( β⒮/β⒟ ) * 1.0296000 * 0.9684867

*/
