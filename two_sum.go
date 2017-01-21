package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var rtn [2]int
	rtn[0] = 0
	found := false
	for !found {
		for idx, num := range nums[rtn[0]+1] {
			if nums[rtn[0]]+num == target {
				rtn[1] = idx
				found = true
				break
			}
		}
		rtn[0] += 1
		if rtn[0] == len(nums)-1 {
			// shoud not ever reach here ...
			break
		}
	}

	return rtn
}

func main() {
	given_nums := int[4]{2, 7, 11, 15}
	target := 9
	fmt.Printf("hello world\n")
}
