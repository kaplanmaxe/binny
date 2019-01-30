package calc

import "testing"

func TestTwosComplementBinary(t *testing.T) {
	answers := make(map[string]string)
	// TODO: add more tests
	answers = map[string]string{
		"00000001": "11111111 (two's complement)",
		"01000001": "10111111 (two's complement)",
		"00000010": "11111110 (two's complement)",
	}

	for unsigned, expected := range answers {
		got, _ := TwosComplementBinary(unsigned)
		if expected != got {
			t.Error(
				"For", unsigned,
				"Expected", expected,
				"Got", got,
			)
		}
	}
}

// func TestTwosComplementDecimal(t *testing.T) {
// 	answers := make(map[int]string)
// 	// TODO: add more tests
// 	answers = map[int]string{
// 		10:  "00001010",
// 		1:   "11111111",
// 		-50: "11001100",
// 	}
// 	for dec, expected := range answers {
// 		got, _ := TwosComplementDecimal(dec)
// 		if got != expected {
// 			t.Error(
// 				"For", dec,
// 				"Expected", expected,
// 				"Got", got,
// 			)
// 		}
// 	}
// }
