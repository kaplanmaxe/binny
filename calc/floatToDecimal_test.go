package calc

import "testing"

func TestFloatToDecimal(t *testing.T) {
	answers := make(map[string]string)
	// TODO: add more tests
	answers = map[string]string{
		"10110100010101100000000000000000": "-1.671875 * 2^-23",
		"01000010011101110000000000000000": "1.929688 * 2^5",
		"11100111010110110000000000000000": "-1.710938 * 2^79",
	}

	for bin, expected := range answers {
		got, _ := FloatToDecimal(bin)
		if got != expected {
			t.Error(
				"For", bin,
				"Expected", expected,
				"Got", got,
			)
		}
	}
}

// TODO: add test for calculateMantisaValue
