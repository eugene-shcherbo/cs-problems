package leetcode

import "github.com/eugene-shcherbo/cs-problems/utils"

func TotalFruits(fruits []int) int {
	bestTotal := 0
	curr := 0
	baskets := make(map[int]int)

	for s, e := 0, 0; e < len(fruits); e++ {
		curr++
		baskets[fruits[e]]++

		for len(baskets) > 2 {
			baskets[fruits[s]]--
			if baskets[fruits[s]] == 0 {
				delete(baskets, fruits[s])
			}
			curr--
		}

		bestTotal = utils.Max(bestTotal, curr)
	}

	return bestTotal
}
