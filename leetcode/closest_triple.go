package leetcode

import (
	"sort"

	"github.com/eugene-shcherbo/cs-problems/utils"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	bestSum := nums[0] + nums[1] + nums[2]

	for i, v := range nums {
		for l, r := i+1, len(nums)-1; l < r; {
			sum := nums[l] + nums[r] + v

			if utils.AbsInt(sum-target) < utils.AbsInt(bestSum-target) {
				bestSum = sum
			}

			if sum > target {
				r--
			} else {
				l++
			}
		}
	}

	return bestSum
}
