package utils

import (
	"crypto/md5"
	"fmt"
	"math"
)

var ColorCache map[string]string = make(map[string]string)

// Converts a PS username into a color in RGB
// Original algorithm source:
// https://github.com/shrianshChari/pokemon-showdown-client/blob/master/play.pokemonshowdown.com/src/battle-log.ts#L562
func UsernameToColor(username string) string {
	id := ToID(username)

	cached, ok := ColorCache[id]
	if ok {
		return cached
	}

	if len(CustomColors) == 0 {
		initializeCustomColors()
	}

	val, ok := CustomColors[id]
	if ok {
		id = ToID(val)
	}

	hash := md5.Sum([]byte(id))

	h := float64((int(hash[2]) * 256 + int(hash[3])) % 360)
	s := float64((int(hash[0]) * 256 + int(hash[1])) % 50 + 40)
	l := float64((int(hash[4]) * 256 + int(hash[5])) % 20 + 30)

	r, g, b := HSLtoRGB(h, s, l)

	lum := r * r * r * 0.2126 + g * g * g * 0.7152 + b * b * b * 0.0722

	HLMod := (lum - 0.2) * -150
	if HLMod > 18 {
		HLMod = (HLMod - 18) * 2.5
	} else if HLMod < 0 {
		HLMod /= 3
	} else {
		HLMod = 0
	}

	HDist := math.Min(math.Abs(180 - h), math.Abs(240 - h))
	if HDist < 15 {
		HLMod += (15 - HDist) / 3
	}

	l += HLMod

	r, g, b = HSLtoRGB(h, s, l)

	r = math.Round(255 * r)
	g = math.Round(255 * g)
	b = math.Round(255 * b)

	color := fmt.Sprintf("#%02x%02x%02x", int(r), int(g), int(b))

	ColorCache[id] = color
	
	return color
}

// Converts a color in HSL to RGB.
// h should be in the range [0, 360]
// s should be in the range [0, 100]
// l should be in the range [0, 100]
//
// This function returns three values in the range [0, 1].
// Multiply these values by 255 to convert them to their color values.
func HSLtoRGB(h float64, s float64, l float64) (float64, float64, float64) {
	chroma := (100 - math.Abs(2 * l - 100)) * s / 100 / 100

	hPrime := h / 60
	x := chroma * (1 - math.Abs(math.Mod(hPrime, 2) - 1))

	m := l / 100 - chroma / 2

	var r, g, b float64;

	switch math.Floor(hPrime) {
	case 0:
		r, g, b = chroma, x, 0
	case 1:
		r, g, b = x, chroma, 0
	case 2:
		r, g, b = 0, chroma, x
	case 3:
		r, g, b = 0, x, chroma
	case 4:
		r, g, b = x, 0, chroma
	default:
		r, g, b = chroma, 0, x
	}

	r = (r + m)
	g = (g + m)
	b = (b + m)

	return r, g, b
}
