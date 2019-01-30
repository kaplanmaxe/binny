package calc

import "testing"

func TestBinToDecUnsigned(t *testing.T) {
	answers := make(map[string]int)
	answers = map[string]int{
		"00000001": 1,
		"11111111": 255,
		"00001010": 10,
		"10001000": 136,
		"11000000": 192,
	}
	for bin, expected := range answers {
		got, _ := BinToDec(bin, false)
		if got != expected {
			t.Error(
				"For", bin,
				"Expected", expected,
				"Got", got,
			)
		}
	}
}

func TestBinToDecSigned(t *testing.T) {
	answers := make(map[string]int)
	// TODO: add more tests
	answers = map[string]int{
		"11111111": -1,
		"11111110": -2,
		"11111000": -8,
	}
	for bin, expected := range answers {
		got, _ := BinToDec(bin, true)
		if got != expected {
			t.Error(
				"For", bin,
				"Expected", expected,
				"Got", got,
			)
		}
	}
}
