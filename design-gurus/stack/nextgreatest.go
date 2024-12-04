package stack

func NextGreatest(nums []int) []int {
	res := make([]int, len(nums))
	candidates := make([]int, 0)

	for i := len(res) - 1; i >= 0; i-- {
		for len(candidates) > 0 && nums[i] >= candidates[len(candidates)-1] {
			candidates = candidates[:len(candidates)-1]
		}

		if len(candidates) > 0 {
			res[i] = candidates[len(candidates)-1]
		} else {
			res[i] = -1
		}

		candidates = append(candidates, nums[i])
	}

	return res
}
