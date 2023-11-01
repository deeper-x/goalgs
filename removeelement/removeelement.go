package removeelement

import "log"

// Solution todo function
func Solution(nums []int, val int) int {
	var tot int
	var others = []int{}

	for _, v := range nums {
		if v == val {
			tot++
		} else {
			others = append(others, v)
		}
	}

	nums = others
	log.Println(nums)
	log.Println(len(nums) - tot)
	return len(nums) - tot
}
