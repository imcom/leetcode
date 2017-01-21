package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	rtn := make([]int, 2, 2)
	rtn[0] = -1
	found := false
	for !found {
		rtn[0] += 1
		for idx, num := range nums[rtn[0]+1:] {
			if nums[rtn[0]]+num == target {
				// need to add offset which is rtn[0]+1
				rtn[1] = idx + rtn[0] + 1
				found = true
				// assume only one solution
				break
			}
		}

		if rtn[0] == len(nums)-1 {
			// shoud not ever reach here ...
			fmt.Printf("oops something went wrong %d:%d\n", rtn[0], len(nums)-1)
			break
		}
	}

	return rtn
}

func main() {
	given_nums := []int{3, 2, 4}
	target := 6
	rtn := twoSum(given_nums, target)
	fmt.Printf("result: %v\n", rtn)
}
