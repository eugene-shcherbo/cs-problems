package greedy

import (
	"sort"
)

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)

	return allTaller(redShirtHeights, blueShirtHeights) || allTaller(blueShirtHeights, redShirtHeights)
}

func allTaller(row1 []int, row2 []int) bool {
	for i, h := range row1 {
		if h <= row2[i] {
			return false
		}
	}
	return true
}
