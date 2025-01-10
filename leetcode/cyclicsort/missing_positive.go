package cyclicsort

import "github.com/eugene-shcherbo/cs-problems/utils"

func FirstMissingPositiveCyclicSort(nums []int) int {
	for i, v := range nums {
		for i+1 != v && v > 0 && v < len(nums) && nums[v-1] != v {
			nums[v-1], nums[i] = nums[i], nums[v-1]
			v = nums[i]
		}
	}

	for i, v := range nums {
		if i+1 != v {
			return i + 1
		}
	}

	return len(nums) + 1
}

func FirstMissingPositiveNegation(nums []int) int {
	for i, v := range nums {
		if v <= 0 {
			nums[i] = len(nums) + 1
		}
	}

	for _, v := range nums {
		v = utils.AbsInt(v)

		if v > len(nums) {
			continue
		}

		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}

	for i, v := range nums {
		if v > 0 {
			return i + 1
		}
	}

	return len(nums) + 1
}
