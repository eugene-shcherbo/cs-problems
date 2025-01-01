package twopointers

import "sort"

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	quadr := make([][]int, 0)

	for i, v := range nums {
		if i == 0 || nums[i] != nums[i-1] {
			quadr = append(quadr, threeSum(nums[i+1:], target-v, v)...)
		}
	}

	return quadr
}

func threeSum(nums []int, target int, firstVal int) [][]int {
	quadr := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		for l, r := i+1, len(nums)-1; l < r; {
			for l != i+1 && l < len(nums) && nums[l] == nums[l-1] {
				l++
			}

			if l >= r {
				break
			}

			sum := nums[i] + nums[l] + nums[r]

			if sum == target {
				quadr = append(quadr, []int{firstVal, nums[i], nums[l], nums[r]})
				l++
				r--
			} else if sum > target {
				r--
			} else {
				l++
			}
		}
	}

	return quadr
}
