package leetcode

import "sort"

func MinimumBoxes(apple []int, capacity []int) int {
	sort.Ints(apple)
	sort.Ints(capacity)

	pack := len(apple) - 1
	for i := len(capacity) - 1; i >= 0; i-- {
		for pack >= 0 && capacity[i] >= apple[pack] {
			capacity[i] -= apple[pack]
			pack--
		}

		if pack < 0 {
			return len(capacity) - i
		}

		apple[pack] -= capacity[i]
	}

	return 0
}
