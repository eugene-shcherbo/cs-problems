package leetcode

import "github.com/eugene-shcherbo/cs-problems/utils"

func SortedSquares(nums []int) []int {
	result := make([]int, 0)

	for l, r := 0, len(nums)-1; l <= r; {
		if nums[l]*nums[l] < nums[r]*nums[r] {
			result = append(result, nums[r]*nums[r])
			r--

		} else {
			result = append(result, nums[l]*nums[l])
			l++
		}
	}

	return utils.Reverse(result)
}
