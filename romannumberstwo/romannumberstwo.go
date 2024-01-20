package romannumberstwo

import (
	"fmt"
)

// Solution returns int number for a given input roman string
func Solution(inVal string) int {
	res := 0

	couples := []string{"IV", "IX", "CD", "CM", "XL", "XC"}
	values := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	for i := 1; i <= len(inVal); i++ {
		first := fmt.Sprintf("%c", inVal[i-1])
		curVal := first
		if i == len(inVal) {
			res += values[curVal]
			return res
		}
		second := fmt.Sprintf("%c", inVal[i])

		couple := fmt.Sprintf("%s%s", first, second)
		isCouple := false

		for y := 0; y < len(couples); y++ {
			if couple == couples[y] {
				isCouple = true
				i++
				break
			}
		}

		if isCouple {
			curVal = couple
		}

		res += values[curVal]
	}

	return res
}
