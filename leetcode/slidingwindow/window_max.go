package slidingwindow

func MaxSlidingWindow(nums []int, k int) []int {
	candidates := make([]int, 0, k)
	maxes := make([]int, 0)

	for i, el := range nums {
		if len(candidates) > 0 && candidates[0] <= i-k {
			candidates = candidates[1:]
		}

		for len(candidates) > 0 && el >= nums[candidates[len(candidates)-1]] {
			candidates = candidates[:len(candidates)-1]
		}

		candidates = append(candidates, i)

		if i >= k-1 {
			maxes = append(maxes, nums[candidates[0]])
		}
	}

	return maxes
}
