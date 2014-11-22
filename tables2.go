package heyday

/*

	// use this

	http://nbviewer.ipython.org/github/colour-science/colour-ipython/blob/master/notebooks/colorimetry/illuminants.ipynb

	// see references and dataset of page above

	2d F11 values from http://en.wikipedia.org/wiki/Standard_illuminant#White_points_of_standard_illuminants

	2 d           x                        y

	A			0.44757					0.40745				2856	Incandescent / Tungsten
	B			0.34842					0.35161				4874	{obsolete} Direct sunlight at noon
	C			0.31006					0.31616				6774	{obsolete} Average / North sky Daylight
	D50			0.34567					0.3585				5003	Horizon Light. ICC profile PCS
	D55			0.33242					0.34743				5503	Mid-morning / Mid-afternoon Daylight
	D60			0.32168					0.33767
	D65			0.31271					0.32902				6504	Noon Daylight: Television, sRGB color space
	D75			0.29902					0.31485				7504	North sky Daylight
	E			0.3333333333333333		0.3333333333333333	5454	(1/3) Equal energy
	F1			0.3131					0.33727				6430	Daylight Fluorescent
	F2			0.37208					0.37529				4230	Cool White Fluorescent
	F3			0.4091					0.3943				3450	White Fluorescent
	F4			0.44018					0.40329				2940	Warm White Fluorescent
	F5			0.31379					0.34531				6350	Daylight Fluorescent
	F6			0.3779					0.38835				4150	Lite White Fluorescent
	F7			0.31292					0.32933				6500	D65 simulator, Daylight simulator
	F8			0.34588					0.35875				5000	D50 simulator, Sylvania F40 Design 50
	F9			0.37417					0.37281				4150	Cool White Deluxe Fluorescent
	F10			0.34609					0.35986				5000	Philips TL85, Ultralume 50
	F11			0.38052					0.37713				4000	Philips TL84, Ultralume 40
	F12			0.43695					0.40441				3000	Philips TL83, Ultralume 30
	FL3.1		0.44067505367083887		0.4032982400521268			New set of fluorescent lamps
	FL3.10		0.345504278345313		0.3559512432054011			New set of fluorescent lamps
	FL3.11		0.3245099692216317		0.343369788104288			New set of fluorescent lamps
	FL3.12		0.4376746480519804		0.40366659242713887			New set of fluorescent lamps
	FL3.13		0.38305207018919446		0.37244321884144194			New set of fluorescent lamps
	FL3.14		0.3447222271025105		0.3609370209728111			New set of fluorescent lamps
	FL3.15		0.3126706172769415		0.3287545660713704			New set of fluorescent lamps
	FL3.2		0.3807750955454208		0.37335593850329946			New set of fluorescent lamps
	FL3.3		0.315282999959924		0.3438607529477053			New set of fluorescent lamps
	FL3.4		0.4429092105157883		0.40432363641263036			New set of fluorescent lamps
	FL3.5		0.37489860509162237		0.36715721144008995			New set of fluorescent lamps
	FL3.6		0.34880430112004895		0.36000066745930004			New set of fluorescent lamps
	FL3.7		0.4384254677714265		0.4045310539294221			New set of fluorescent lamps
	FL3.8		0.38197740008964315		0.3831793433643612			New set of fluorescent lamps
	FL3.9		0.3498569555988198		0.3590880129881551			New set of fluorescent lamps
	HP1			0.5330008227923825		0.4149532376869364			High pressure discharge lamps
	HP2			0.4777920749861397		0.4158404588084794			High pressure discharge lamps
	HP3			0.4302314003272559		0.4075170114048384			High pressure discharge lamps
	HP4			0.38117189471910373		0.3797265745279147			High pressure discharge lamps
	HP5			0.37758320909195225		0.3713479728022603			High pressure discharge lamps


	10d					x						y

	A				0.45117					0.40594
	B				0.3498					0.3527
	C				0.31039					0.31905
	D50				0.34773					0.35952
	D55				0.33411					0.34877
	D60				0.3229915227773675		0.3391283129096501
	D65				0.31382					0.331
	D75				0.29968					0.3174
	E				0.3333333333333333		0.3333333333333333			(1/3)
	F1				0.31811					0.33559
	F10				0.3509					0.35444
	F11				0.38541					0.37123
	F12				0.44256					0.39717
	F2				0.37925					0.36733
	F3				0.41761					0.38324
	F4				0.4492					0.39074
	F5				0.31975					0.34246
	F6				0.3866					0.37847
	F7				0.31569					0.3296
	F8				0.34902					0.35939
	F9				0.37829					0.37045
	FL3.1			0.44983072060097606		0.3902314548749185
	FL3.10			0.3493448005253757		0.35498443751900227
	FL3.11			0.32926800293079544		0.33886544553697306
	FL3.12			0.4422521047116721		0.40122058660685506
	FL3.13			0.3862752927999335		0.3742832325728946
	FL3.14			0.34725510277403043		0.3668082930976077
	FL3.15			0.31461426562662165		0.3333777810577307
	FL3.2			0.3869241790592931		0.3657561399067915
	FL3.3			0.3211770463590012		0.3405012294853132
	FL3.4			0.44812132825245227		0.39707718739452286
	FL3.5			0.3778142145280147		0.36662585171298506
	FL3.6			0.35197652659197876		0.3610945283215212
	FL3.7			0.44430926608426524		0.3967914695376695
	FL3.8			0.3875889785058527		0.37630564994236454
	FL3.9			0.3546890235509348		0.3534450969379135
	HP1				0.543334641014127		0.40528934193733374
	HP2				0.48264737821938364		0.4108157059894971
	HP3				0.4355600804884421		0.39880115305546204
	HP4				0.3851937462655714		0.3682756593039444
	HP5				0.38031641910835223		0.3666171209128104

	White point

	Y = 100

	X = Yx/y

	Z = Y(1-x-y)/y

*/
