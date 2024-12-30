package leetcode

import "github.com/eugene-shcherbo/cs-problems/utils"

func LengthOfLongestSubstringKDistinct(s string, k int) int {
	distinct := make(map[byte]int)
	bestSize := 0

	for l, r := 0, 0; r < len(s); r++ {
		distinct[s[r]]++

		for l <= r && len(distinct) > k {
			distinct[s[l]]--
			if distinct[s[l]] == 0 {
				delete(distinct, s[l])
			}
			l++
		}

		bestSize = utils.Max(bestSize, r-l+1)
	}

	return bestSize
}
