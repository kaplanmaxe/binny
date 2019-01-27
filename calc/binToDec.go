package calc

import (
	"math"
	"strconv"
)

// BinToDec converts a binary number to decimal
func BinToDec(bin string, signed bool) (int, error) {
	digits := []rune(bin)
	pow := 7
	dec := 0
	for key, digit := range digits {
		d, err := strconv.Atoi(string(digit))
		if err != nil {
			return 0, err
		}

		if signed == true && (key == 0 && d == 1) {
			dec += int((float64(d) * -1) * math.Pow(2, float64(pow)))
		} else {
			dec += int(float64(d) * math.Pow(2, float64(pow)))
		}
		pow--
	}
	return dec, nil
}
