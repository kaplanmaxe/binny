package calc

import (
	"errors"
	"fmt"
	"math"
)

const bias = 127

func FloatToDecimal(bin string) (string, error) {
	if len(bin) != 32 {
		return "", errors.New("Bit string must be 32 bits long")
	}
	bits := []rune(bin)
	signBit := bits[0]
	var sign int
	if signBit == '0' {
		sign = 1
	} else {
		sign = -1
	}
	exponent, err := BinToDec(string(bits[1:9]), false)
	if err != nil {
		return "", err
	}
	exponent -= bias
	mantisaBits := bits[9:32]
	mantisa := calculateMantisaValue(mantisaBits)
	fraction := mantisa * float64(sign)
	return fmt.Sprintf("%f * 2^%d", fraction, exponent), nil
}

func calculateMantisaValue(bits []rune) float64 {
	const offset = -1
	var result float64
	for i := 0; i < 23; i++ {
		exponent := float64(0 - (i + 1))
		if bits[i] == '1' {
			result += math.Pow(2.0, exponent)
		}
	}
	return result + 1
}
