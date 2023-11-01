package romanumbers

import (
	"fmt"
	"log"
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
	res := []string{}

	log.Println(s)
	for i := 0; i < len(tokens); i++ {

		isCouple := false

		if i < lt-1 {
			//log.Println(tokens[first], tokens[second])

			couple := fmt.Sprintf("%s%s", tokens[i], tokens[i+1])
			log.Println(couple)
			for _, v1 := range subt {
				if couple == v1 {
					log.Println(couple, " is a couple")
					res = append(res, couple)
					i++
					isCouple = true
					break
				}
			}
		}

		if isCouple {
			continue
		}

		log.Println(tokens[i], " is a single")
		res = append(res, tokens[i])

		log.Println(i, i+1)
	}

	log.Println(res)

	tot := 0
	for _, v := range res {
		tot += getVal(v)
	}

	return tot
}
