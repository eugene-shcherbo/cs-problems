package leetcode

import (
	"math"
	"sort"
)

func FindLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	maxLen := 0
	last := math.MinInt

	for _, pair := range pairs {
		if pair[0] > last {
			maxLen++
			last = pair[1]
		}
	}

	return maxLen
}
