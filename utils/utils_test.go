package utils

import (
	"math"
	"reflect"
	"testing"
)

type toIdTest struct {
	input, expected string
}

var toIdTests = []toIdTest {
	{"abcd", "abcd"},
	{"Abcd", "abcd"},
	{"ABcd", "abcd"},
	{"ABCd", "abcd"},
	{"ABCD", "abcd"},
	{"ab cd", "abcd"},
	{"abcd!", "abcd"},
	{"ab cd!", "abcd"},
	{"AB CD!", "abcd"},
	{"Ab_cd?", "abcd"},
	{"AB-cd.", "abcd"},
}

func TestToID(t *testing.T) {
	for _, test := range toIdTests {
		output := ToID(test.input)

		if (output != test.expected) {
			t.Fatalf("Input: %v, Expected: %v, Actual: %v", test.input, test.expected, output)
		}
	}
}

type hslTest struct {
	h, s, l, expectedR, expectedG, expectedB float64
}

var hslTests = []hslTest{
	{0, 0, 0, 0, 0, 0},
	{64, 36, 33, 110, 114, 54},
	{83, 87, 3, 9, 14, 1},
	{271, 74, 83, 213, 180, 244},
	{281, 66, 49, 155, 42, 207},
}

func TestHSLtoRGB(t *testing.T) {
	for _, test := range hslTests {
		r, g, b := HSLtoRGB(test.h, test.s, test.l)
		r = math.Round(r * 255)
		g = math.Round(g * 255)
		b = math.Round(b * 255)

		input := [3]float64{test.h, test.s, test.l}
		want := [3]float64{test.expectedR, test.expectedG, test.expectedB}
		output := [3]float64{r, g, b}
		if (output != want) {
			t.Fatalf("Input: %v, Expected: %v, Actual: %v", input, want, output)
		}
	}
}

var usernameColorTests = map[string]string{
	"Tb0lt": "#ab5412",
	"a": "#8162d1",
	"b": "#66a535",
	"c": "#1aa22d",
	"d": "#5539f3",
	"e": "#924db8",
	// Tests usage of custom colors
	"justusegoogle": "#0cacba",
	"dhelmise": "#0a8ca0",
	"chiyu": "#df6a11",
	"snivy": "#33a726",
	"kyogre": "#2972df",
}

func TestUsernameToColor(t *testing.T) {
	for username, expected := range usernameColorTests {
		output := UsernameToColor(username)

		if (output != expected) {
			t.Fatalf("Input: %v, Expected: %v, Actual: %v", username,
				expected, output)
		}
	}
}

func TestColorCache(t *testing.T) {
	if reflect.DeepEqual(usernameColorTests, ColorCache) {
		t.Fatalf("ColorCache: Expected: %v, Actual: %v", usernameColorTests, ColorCache)
	}
}
