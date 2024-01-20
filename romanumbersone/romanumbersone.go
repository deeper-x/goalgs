package romanumbersone

import (
	"fmt"
	"strings"
)

var subt = []string{"IV", "IX", "XL", "XC", "CD", "CM"}
var valMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func getVal(inVal string) int {
	sliceVals := strings.Split(inVal, "")

	if len(sliceVals) == 2 {
		return valMap[sliceVals[1]] - valMap[sliceVals[0]]
	}

	return valMap[sliceVals[0]]
}

// Solution roman numbers
func Solution(s string) int {
	tokens := strings.Split(s, "")
	lt := len(tokens)
	tot := 0

	for i := 0; i < lt; i++ {
		isCouple := false

		if i < lt-1 {
			couple := fmt.Sprintf("%s%s", tokens[i], tokens[i+1])
			for _, v1 := range subt {
				if couple == v1 {
					tot += getVal(couple)
					i++
					isCouple = true
					break
				}
			}
		}

		if isCouple {
			continue
		}

		tot += getVal(tokens[i])
	}

	return tot
}
