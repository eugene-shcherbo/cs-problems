package leetcode

import "sort"

func DivideArray(nums []int, k int) [][]int {
	sort.Ints(nums)

	nGroups := len(nums) / 3
	res := make([][]int, 0)

	for i := 0; i < nGroups; i++ {
		if nums[i*3+2]-nums[i*3] > k {
			return nil
		}
		res = append(res, nums[i*3:i*3+3])
	}

	return res
}
