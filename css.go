package heyday

import (
	"errors"
	"strconv"
	"strings"
)

func Css(s string) (*Color, error) {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	switch {
	case strings.HasPrefix(s, "#"):
		return FromHex(s)
	default:
		return FromCss(s)
	}
}

func FromHex(s) (*Color, error) {
	s = strings.Replace(s, " ", "", -1)
	s = s[1:] // #123 => 123
	var err error
	var r, g, b, a int64
	var alpha bool
	switch len(s) {
	case 6: // RRGGBB
		r, err = strconv.ParseInt(s[:2], 16, 64)
		g, err = strconv.ParseInt(s[2:4], 16, 64)
		b, err = strconv.ParseInt(s[4:], 16, 64)
	case 3: // RGB
		r, err = strconv.ParseInt(s[:1], 16, 64)
		g, err = strconv.ParseInt(s[1:2], 16, 64)
		b, err = strconv.ParseInt(s[2:], 16, 64)
	case 8: // AARRGGBB // like qml
		alpha = true
		a, err = strconv.ParseInt(s[:2], 16, 64)
		r, err = strconv.ParseInt(s[2:4], 16, 64)
		g, err = strconv.ParseInt(s[4:6], 16, 64)
		b, err = strconv.ParseInt(s[6:], 16, 64)
	case 4: // ARGB
		alpha = true
		a, err = strconv.ParseInt(s[:1], 16, 64)
		r, err = strconv.ParseInt(s[1:2], 16, 64)
		g, err = strconv.ParseInt(s[2:3], 16, 64)
		b, err = strconv.ParseInt(s[3:], 16, 64)
	default:
		return nil, errors.New("Couldn't detect hex color")
	}
	if err != nil {
		return nil, err
	}
	if alpha {
		return &Color(RGBA{float64(r), float64(g), float64(b), float64(a)}), nil
	} else {
		return &Color(RGB{float64(r), float64(g), float64(b)}), nil
	}
}

func FromCss(s string) (*Color, error) {
	switch {
	case strings.HasPrefix(s, "rgba"):
		if r, g, b, a, err := decode_css_string_with_alpha(s); err == nil {
			return &Color(RGBA{r, g, b, a}), nil
		} else {
			return nil, err
		}
	case strings.HasPrefix(s, "hsla"):
		if h, s, l, a, err := decode_css_string_with_alpha(s); err == nil {
			return &Color(HSLA{h, s, l, a}), nil
		} else {
			return nil, err
		}
	case strings.HasPrefix(s, "rgb"):
		if r, g, b, err := decode_css_string(s); err == nil {
			return &Color(RGB{r, g, b}), nil
		} else {
			return nil, err
		}
	case strings.HasPrefix(s, "hsl"):
		if h, s, l, err := decode_css_string(s); err == nil {
			return &Color(HSL{h, s, l}), nil
		} else {
			return nil, err
		}
	case strings.HasPrefix(s, "hsv") || strings.HasPrefix(s, "hsb"):
		if h, s, v, err := decode_css_string(s); err == nil {
			return &Color(HSV{h, s, v}), nil
		} else {
			return nil, err
		}
	case strings.HasPrefix(s, "cmyk"):
		if c, m, y, k, err := decode_cmyk_css_string(s); err == nil {
			return &Color(CMYK{c, m, y, k}), nil
		} else {
			return nil, err
		}
	case strings.HasPrefix(s, "cmy"):
		if c, m, y, err := decode_css_string(s); err == nil {
			return &Color(CMY{c, m, y}), nil
		} else {
			return nil, err
		}
	default:
		return nil, errors.New("Couldn't detect css color")
	}
}

func decode_css_string(s string) (r, g, b float64) {
	//
}

func decode_css_string_with_alpha(s string) (r, g, b, a float64) {
	//
}

func decode_cmyk_css_string(s string) (c, m, y, k float64) {
	//
}
