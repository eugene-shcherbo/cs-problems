package arrays

import "sort"

func ContainsDuplicatesSort(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

func ContainsDuplicatesHashTable(nums []int) bool {
	unique := make(map[int]bool)

	for _, num := range nums {
		if unique[num] {
			return true
		}
		unique[num] = true
	}

	return false
}
