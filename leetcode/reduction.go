package leetcode

import "sort"

func ReductionOperations(nums []int) int {
	sort.Ints(nums)

	ops := 0
	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			idx++
		}
		ops += idx
	}

	return ops
}
