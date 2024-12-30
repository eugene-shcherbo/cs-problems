package leetcode

import "github.com/eugene-shcherbo/cs-problems/utils"

func LongestOnes(nums []int, k int) int {
	zeroes := 0
	bestLen := 0

	for s, e := 0, 0; e < len(nums); e++ {
		if nums[e] == 0 {
			zeroes++
		}

		for zeroes > k {
			if nums[s] == 0 {
				zeroes--
			}
			s++
		}

		bestLen = utils.Max(bestLen, e-s+1)
	}

	return bestLen
}
