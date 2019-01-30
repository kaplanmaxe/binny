package calc

import "testing"

func TestDecToBin(t *testing.T) {
	answers := make(map[int]string)
	answers = map[int]string{
		128: "10000000",
		10:  "00001010",
		192: "11000000",
		1:   "00000001",
		130: "10000010",
		5:   "00000101",
	}

	for dec, expected := range answers {
		got, _ := DecToBin(dec)
		if got != expected {
			t.Error(
				"For", dec,
				"Expected", expected,
				"Got", got,
			)
		}
	}
}
