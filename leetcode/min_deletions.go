package leetcode

import "github.com/eugene-shcherbo/cs-problems/utils"

func MinimumDeletions(nums []int) int {
	minPos := 0
	maxPos := 0

	for i, v := range nums {
		if v > nums[maxPos] {
			maxPos = i
		}
		if v < nums[minPos] {
			minPos = i
		}
	}

	bigger := minPos
	smaller := maxPos
	if bigger < smaller {
		bigger, smaller = smaller, bigger
	}

	return utils.Min(bigger+1, utils.Min(len(nums)-smaller, smaller+1+len(nums)-bigger))
}
