package heaps

import "github.com/eugene-shcherbo/cs-problems/ds"

func ConnectSticks(sticks []int) int {
	heap := ds.NewHeapComp[int](sticks)

	totalCost := 0

	for heap.Len() > 1 {
		first, _ := heap.Extract()
		second, _ := heap.Extract()
		heap.Add(first + second)

		totalCost += first + second
	}

	return totalCost
}
