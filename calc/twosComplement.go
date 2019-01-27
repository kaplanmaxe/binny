package calc

import (
	"fmt"
)

// TwosComplementBinary calculate the two's complement binary string
func TwosComplementBinary(bin string) (string, error) {
	bits := []rune(bin)
	var nBits []rune
	for _, bit := range bits {
		if bit == '1' {
			nBits = append(nBits, '0')
		} else if bit == '0' {
			nBits = append(nBits, '1')
		}
	}

	// TODO: do this with binary arithmetic
	dec, err := BinToDec(string(nBits), false)
	if err != nil {
		return "00000000", err
	}
	// TODO: do this with binary arithmetic
	twosComplement, err := DecToBin(dec + 1)
	return fmt.Sprintf("%s (two's complement)", twosComplement), err
}

// TwosComplementDecimal takes a number in decimal and returns it's two's complement
func TwosComplementDecimal(n int) (string, error) {
	var bin string
	var unsignedBinary string
	var err error
	if n > 0 {
		bin, err = DecToBin(n)
	} else {
		unsignedBinary, err = DecToBin(n)
		bin, err = TwosComplementBinary(unsignedBinary)
	}
	return bin, err
}
