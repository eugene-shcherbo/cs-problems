package intervals

import "github.com/eugene-shcherbo/cs-problems/utils"

func Insert(intervals [][]int, newInterval []int) [][]int {
	newIntervals := make([][]int, 0)

	for _, interval := range intervals {
		if interval[0] > newInterval[0] {
			break
		}
		newIntervals = append(newIntervals, interval)
	}

	if len(newIntervals) == 0 || !merge(newIntervals[len(newIntervals)-1], newInterval) {
		newIntervals = append(newIntervals, newInterval)
	}

	for i := len(newIntervals) - 1; i < len(intervals); i++ {
		if !merge(newIntervals[len(newIntervals)-1], intervals[i]) {
			newIntervals = append(newIntervals, intervals[i])
		}
	}

	return newIntervals
}

func merge(left, right []int) bool {
	if left[1] >= right[0] {
		left[1] = utils.Max(left[1], right[1])
		return true
	}
	return false
}
