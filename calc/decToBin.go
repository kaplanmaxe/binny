package calc

import (
	"errors"
	"math"
)

// DecToBin takes a number in decimal and converts to binary
func DecToBin(n int) (string, error) {
	var negative bool
	if n > 255 {
		return "00000000", errors.New("Numbers over 255 not yet supported")
	}
	if n < 0 {
		n *= -1
		negative = true
	}

	pow := 7
	remainder := n
	var digits = []rune{}
	for i := pow; i >= 0; i-- {
		binVal := int(math.Pow(2, float64(i)))

		if remainder == 0 && len(digits) > 0 {
			digits = append(digits, '0')
			continue
		}
		if remainder-binVal >= 0 {
			digits = append(digits, '1')
			remainder -= binVal
		} else {
			digits = append(digits, '0')
		}
	}
	if negative {
		return TwosComplementBinary(string(digits))
	}
	return string(digits), nil
}
