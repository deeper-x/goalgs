package main

import (
	"log"

	"github.com/deeper-x/goalgs/romanumbers"
)

func main() {
	// var a []int = []int{3, 2, 2, 3}
	// removeelement.Solution(a, 3)

	res := romanumbers.Solution("MCMXCIV")
	log.Println(res)
}
