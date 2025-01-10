package cyclicsort

// time: O(N)
// space: O(N)
func FindDisappearedNumbersSpace(nums []int) []int {
	allSeen := make([]bool, len(nums))

	for _, n := range nums {
		allSeen[n-1] = true
	}

	missing := make([]int, 0)
	for i, seen := range allSeen {
		if !seen {
			missing = append(missing, i+1)
		}
	}

	return missing
}

// time: O(N)
// space: O(1)
func FindDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		correct_idx := nums[i] - 1
		if correct_idx != i && nums[i] != nums[correct_idx] {
			nums[i], nums[correct_idx] = nums[correct_idx], nums[i]
		} else {
			i++
		}
	}

	missing := make([]int, 0)
	for i, val := range nums {
		if i != val-1 {
			missing = append(missing, i+1)
		}
	}

	return missing
}
