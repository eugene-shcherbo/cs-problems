package cyclicsort

import "github.com/eugene-shcherbo/cs-problems/utils"

func FindDuplicatesCyclicSort(nums []int) []int {
	for i, v := range nums {
		for nums[v-1] != nums[i] {
			nums[v-1], nums[i] = nums[i], nums[v-1]
			v = nums[i]
		}
	}

	duplicates := make([]int, 0)
	for i, v := range nums {
		if i+1 != v {
			duplicates = append(duplicates, v)
		}
	}

	return duplicates
}

func FindDuplicatesNegation(nums []int) []int {
	duplicates := make([]int, 0)

	for _, n := range nums {
		correctPos := utils.AbsInt(n) - 1

		if nums[correctPos] < 0 {
			duplicates = append(duplicates, utils.AbsInt(n))
		} else {
			nums[correctPos] = -nums[correctPos]
		}
	}

	return duplicates
}
