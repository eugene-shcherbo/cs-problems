package leetcode

import "github.com/eugene-shcherbo/cs-problems/utils"

func MinSubArrayLen(target int, nums []int) int {
	sum := 0
	bestSize := len(nums) + 1

	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]

		for ; l <= r && sum >= target; l++ {
			bestSize = utils.Min(bestSize, r-l+1)
			sum -= nums[l]
		}
	}

	if bestSize == len(nums)+1 {
		return 0
	}

	return bestSize
}
