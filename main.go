package main

import (
	"log"

	"github.com/deeper-x/goalgs/removeelement"
	"github.com/deeper-x/goalgs/romannumberstwo"
	"github.com/deeper-x/goalgs/romanumbersone"
)

func main() {
	var a []int = []int{3, 2, 2, 3}
	removeelement.Solution(a, 3)

	res1 := romanumbersone.Solution("MCMXCIV")
	log.Println(res1)

	res2 := romannumberstwo.Solution("MCMXCIV")
	log.Println(res2)

	res3 := romannumberstwo.Solution("III")
	log.Println(res3)

	res4 := romannumberstwo.Solution("MDCXCV")
	log.Println(res4)
}
