package heyday

import (
	"errors"
	"strings"
)

var Named = map[string]RGB{}

/*
Contains named colors
*/

func init() {
	// Pinks
	Named["pink"] = RGB{255.0, 192.0, 203.0}           // Pink
	Named["lightpink"] = RGB{255.0, 182.0, 193.0}      // Light Pink
	Named["hotpink"] = RGB{255.0, 105.0, 180.0}        // Hot Pink
	Named["deeppink"] = RGB{255.0, 20.0, 147.0}        // Deep Pink
	Named["palevioletred"] = RGB{219.0, 112.0, 147.0}  // Pale Violet Red
	Named["mediumvioletred"] = RGB{199.0, 21.0, 133.0} // Medium Violet Red
	// Reds
	Named["lightsalmon"] = RGB{255.0, 160.0, 122.0} // Light Salmon
	Named["salmon"] = RGB{250.0, 128.0, 114.0}      // Salmon
	Named["darksalmon"] = RGB{233.0, 150.0, 122.0}  // Dark Salmon
	Named["lightcoral"] = RGB{240.0, 128.0, 128.0}  // Light Coral
	Named["indianred"] = RGB{205.0, 92.0, 92.0}     // Indian Red
	Named["crimson"] = RGB{220.0, 20.0, 60.0}       // Crimson
	Named["firebrick"] = RGB{178.0, 34.0, 34.0}     // Fire Brick
	Named["darkred"] = RGB{139.0, 0.0, 0.0}         // Dark Red
	Named["red"] = RGB{255.0, 0.0, 0.0}             // Red
	// Oranges
	Named["orangered"] = RGB{255.0, 69.0, 0.0}   // Orange Red
	Named["tomato"] = RGB{255.0, 99.0, 71.0}     // Tomato
	Named["coral"] = RGB{255.0, 127.0, 80.0}     // Coral
	Named["darkorange"] = RGB{255.0, 140.0, 0.0} // Dark Orange
	Named["orange"] = RGB{255.0, 165.0, 0.0}     // Orange
	// Yellows
	Named["yellow"] = RGB{255.0, 255.0, 0.0}                 // Yellow
	Named["lightyellow"] = RGB{255.0, 255.0, 224.0}          // Light Yellow
	Named["lemonchiffon"] = RGB{255.0, 250.0, 205.0}         // Lemon Chiffon
	Named["lightgoldenrodyellow"] = RGB{250.0, 250.0, 210.0} // Light Goldenrod Yellow
	Named["papayawhip"] = RGB{255.0, 239.0, 213.0}           // Papaya Whip
	Named["moccasin"] = RGB{255.0, 228.0, 181.0}             // Moccasin
	Named["peachpuff"] = RGB{255.0, 218.0, 185.0}            // Peach Puff
	Named["palegoldenrod"] = RGB{238.0, 232.0, 170.0}        // Pale Goldenrod
	Named["khaki"] = RGB{240.0, 230.0, 140.0}                // Khaki
	Named["darkkhaki"] = RGB{189.0, 183.0, 107.0}            // Dark Khaki
	Named["gold"] = RGB{255.0, 215.0, 0.0}                   // Gold
	// Browns
	Named["cornsilk"] = RGB{255.0, 248.0, 220.0}       // Cornsilk
	Named["blanchedalmond"] = RGB{255.0, 235.0, 205.0} // Blanched Almond
	Named["bisque"] = RGB{255.0, 228.0, 196.0}         // Bisque
	Named["navajowhite"] = RGB{255.0, 222.0, 173.0}    // Navajo White
	Named["wheat"] = RGB{245.0, 222.0, 179.0}          // Wheat
	Named["burlywood"] = RGB{222.0, 184.0, 135.0}      // Burly Wood
	Named["tan"] = RGB{210.0, 180.0, 140.0}            // Tan
	Named["rosybrown"] = RGB{188.0, 143.0, 143.0}      // Rosy Brown
	Named["sandybrown"] = RGB{244.0, 164.0, 96.0}      // Sandy Brown
	Named["goldenrod"] = RGB{218.0, 165.0, 32.0}       // Goldenrod
	Named["darkgoldenrod"] = RGB{184.0, 134.0, 11.0}   // Dark Goldenrod
	Named["peru"] = RGB{205.0, 133.0, 63.0}            // Peru
	Named["chocolate"] = RGB{210.0, 105.0, 30.0}       // Chocolate
	Named["saddlebrown"] = RGB{139.0, 69.0, 19.0}      // Saddle Brown
	Named["sienna"] = RGB{160.0, 82.0, 45.0}           // Sienna
	Named["brown"] = RGB{165.0, 42.0, 42.0}            // Brown
	Named["maroon"] = RGB{128.0, 0.0, 0.0}             // Maroon
	// Greens
	Named["darkolivegreen"] = RGB{85.0, 107.0, 47.0}    // Dark OliveGreen
	Named["olive"] = RGB{128.0, 128.0, 0.0}             // Olive
	Named["olivedrab"] = RGB{107.0, 142.0, 35.0}        // Olive Drab
	Named["yellowgreen"] = RGB{154.0, 205.0, 50.0}      // Yellow Green
	Named["limegreen"] = RGB{50.0, 205.0, 50.0}         // Lime Green
	Named["lime"] = RGB{0.0, 255.0, 0.0}                // Lime
	Named["lawngreen"] = RGB{124.0, 252.0, 0.0}         // Lawn Green
	Named["chartreuse"] = RGB{127.0, 255.0, 0.0}        // Chartreuse
	Named["greenyellow"] = RGB{173.0, 255.0, 47.0}      // Green Yellow
	Named["springgreen"] = RGB{0.0, 255.0, 127.0}       // Spring Green
	Named["mediumspringgreen"] = RGB{0.0, 250.0, 154.0} // Medium Spring Green
	Named["lightgreen"] = RGB{144.0, 238.0, 144.0}      // Light Green
	Named["palegreen"] = RGB{152.0, 251.0, 152.0}       // Pale Green
	Named["darkseagreen"] = RGB{143.0, 188.0, 143.0}    // Dark Sea Green
	Named["mediumseagreen"] = RGB{60.0, 179.0, 113.0}   // Medium Sea Green
	Named["seagreen"] = RGB{46.0, 139.0, 87.0}          // Sea Green
	Named["forestgreen"] = RGB{34.0, 139.0, 34.0}       // Forest Green
	Named["green"] = RGB{0.0, 128.0, 0.0}               // Green
	Named["darkgreen"] = RGB{0.0, 100.0, 0.0}           // Dark Green
	// Cyans
	Named["mediumaquamarine"] = RGB{102.0, 205.0, 170.0} // Medium Aquamarine
	Named["aqua"] = RGB{0.0, 255.0, 255.0}               // Aqua
	Named["cyan"] = RGB{0.0, 255.0, 255.0}               // Cyan
	Named["lightcyan"] = RGB{224.0, 255.0, 255.0}        // Light Cyan
	Named["paleturquoise"] = RGB{175.0, 238.0, 238.0}    // Pale Turquoise
	Named["aquamarine"] = RGB{127.0, 255.0, 212.0}       // Aquamarine
	Named["turquoise"] = RGB{64.0, 224.0, 208.0}         // Turquoise
	Named["mediumturquoise"] = RGB{72.0, 209.0, 204.0}   // Medium Turquoise
	Named["darkturquoise"] = RGB{0.0, 206.0, 209.0}      // Dark Turquoise
	Named["lightseagreen"] = RGB{32.0, 178.0, 170.0}     // Light Sea Green
	Named["cadetblue"] = RGB{95.0, 158.0, 160.0}         // Cadet Blue
	Named["darkcyan"] = RGB{0.0, 139.0, 139.0}           // Dark Cyan
	Named["teal"] = RGB{0.0, 128.0, 128.0}               // Teal
	// Blues
	Named["lightsteelblue"] = RGB{176.0, 196.0, 222.0} // Light Steel Blue
	Named["powderblue"] = RGB{176.0, 224.0, 230.0}     // Powder Blue
	Named["lightblue"] = RGB{173.0, 216.0, 230.0}      // Light Blue
	Named["skyblue"] = RGB{135.0, 206.0, 235.0}        // Sky Blue
	Named["lightskyblue"] = RGB{135.0, 206.0, 250.0}   // Light SkyBlue
	Named["deepskyblue"] = RGB{0.0, 191.0, 255.0}      // Deep Sky Blue
	Named["dodgerblue"] = RGB{30.0, 144.0, 255.0}      // Dodger Blue
	Named["cornflowerblue"] = RGB{100.0, 149.0, 237.0} // Cornflower Blue
	Named["steelblue"] = RGB{70.0, 130.0, 180.0}       // Steel Blue
	Named["royalblue"] = RGB{65.0, 105.0, 225.0}       // Royal Blue
	Named["blue"] = RGB{0.0, 0.0, 255.0}               // Blue
	Named["mediumblue"] = RGB{0.0, 0.0, 205.0}         // Medium Blue
	Named["darkblue"] = RGB{0.0, 0.0, 139.0}           // Dark Blue
	Named["navy"] = RGB{0.0, 0.0, 128.0}               // Navy
	Named["midnightblue"] = RGB{25.0, 25.0, 112.0}     // Midnight Blue
	// Purples
	Named["lavender"] = RGB{230.0, 230.0, 250.0}        // Lavender
	Named["thistle"] = RGB{216.0, 191.0, 216.0}         // Thistle
	Named["plum"] = RGB{221.0, 160.0, 221.0}            // Plum
	Named["violet"] = RGB{238.0, 130.0, 238.0}          // Violet
	Named["orchid"] = RGB{218.0, 112.0, 214.0}          // Orchid
	Named["fuchsia"] = RGB{255.0, 0.0, 255.0}           // Fuchsia
	Named["magenta"] = RGB{255.0, 0.0, 255.0}           // Magenta
	Named["mediumorchid"] = RGB{186.0, 85.0, 211.0}     // Medium Orchid
	Named["mediumpurple"] = RGB{147.0, 112.0, 219.0}    // Medium Purple
	Named["blueviolet"] = RGB{138.0, 43.0, 226.0}       // Blue Violet
	Named["darkviolet"] = RGB{148.0, 0.0, 211.0}        // Dark Violet
	Named["darkorchid"] = RGB{153.0, 50.0, 204.0}       // Dark Orchid
	Named["darkmagenta"] = RGB{139.0, 0.0, 139.0}       // Dark Magenta
	Named["purple"] = RGB{128.0, 0.0, 128.0}            // Purple
	Named["indigo"] = RGB{75.0, 0.0, 130.0}             // Indigo
	Named["darkslateblue"] = RGB{72.0, 61.0, 139.0}     // Dark Slate Blue
	Named["slateblue"] = RGB{106.0, 90.0, 205.0}        // Slate Blue
	Named["mediumslateblue"] = RGB{123.0, 104.0, 238.0} // Medium Slate Blue
	// Whites
	Named["white"] = RGB{255.0, 255.0, 255.0}         // White
	Named["snow"] = RGB{255.0, 250.0, 250.0}          // Snow
	Named["honeydew"] = RGB{240.0, 255.0, 240.0}      // Honeydew
	Named["mintcream"] = RGB{245.0, 255.0, 250.0}     // Mint Cream
	Named["azure"] = RGB{240.0, 255.0, 255.0}         // Azure
	Named["aliceblue"] = RGB{240.0, 248.0, 255.0}     // Alice Blue
	Named["ghostwhite"] = RGB{248.0, 248.0, 255.0}    // Ghost White
	Named["whitesmoke"] = RGB{245.0, 245.0, 245.0}    // White Smoke
	Named["seashell"] = RGB{255.0, 245.0, 238.0}      // Seashell
	Named["beige"] = RGB{245.0, 245.0, 220.0}         // Beige
	Named["oldlace"] = RGB{253.0, 245.0, 230.0}       // Old Lace
	Named["floralwhite"] = RGB{255.0, 250.0, 240.0}   // Floral White
	Named["ivory"] = RGB{255.0, 255.0, 240.0}         // Ivory
	Named["antiquewhite"] = RGB{250.0, 235.0, 215.0}  // Antique White
	Named["linen"] = RGB{250.0, 240.0, 230.0}         // Linen
	Named["lavenderblush"] = RGB{255.0, 240.0, 245.0} // Lavender Blush
	Named["mistyrose"] = RGB{255.0, 228.0, 225.0}     // Misty Rose
	// Grays/Blacks
	Named["gainsboro"] = RGB{220.0, 220.0, 220.0}      // Gainsboro
	Named["lightgrey"] = RGB{211.0, 211.0, 211.0}      // Light Grey
	Named["silver"] = RGB{192.0, 192.0, 192.0}         // Silver
	Named["darkgray"] = RGB{169.0, 169.0, 169.0}       // Dark Gray
	Named["darkgrey"] = RGB{169.0, 169.0, 169.0}       // Dark Grey
	Named["gray"] = RGB{128.0, 128.0, 128.0}           // Gray
	Named["grey"] = RGB{128.0, 128.0, 128.0}           // Grey
	Named["dimgray"] = RGB{105.0, 105.0, 105.0}        // Dim Gray
	Named["dimgrey"] = RGB{105.0, 105.0, 105.0}        // Dim Grey
	Named["lightslategray"] = RGB{119.0, 136.0, 153.0} // Light Slate Gray
	Named["lightslategrey"] = RGB{119.0, 136.0, 153.0} // Light Slate Grey
	Named["slategray"] = RGB{112.0, 128.0, 144.0}      // Slate Gray
	Named["slategrey"] = RGB{112.0, 128.0, 144.0}      // Slate Grey
	Named["darkslategray"] = RGB{47.0, 79.0, 79.0}     // Dark Slate Gray
	Named["darkslategrey"] = RGB{47.0, 79.0, 79.0}     // Dark Slate Grey
	Named["black"] = RGB{0.0, 0.0, 0.0}                // Black
}

func ByName(name string) (*RGB, error) {
	lower := strings.ToLower(name)
	if clr, ok := Named[lower]; !ok {
		return nil, errors.New("No such named color, name: %s", name)
	} else {
		return &clr, nil
	}
}

/*
Get RGB color by name. Available names is all X11 named colors.
See that http://en.wikipedia.org/wiki/Web_colors#X11_color_names table for
details. All *gray* may be replaced by *grey*. Aqua = Cyan, Magenta = Fuchsia.
Case does not matter
*/
