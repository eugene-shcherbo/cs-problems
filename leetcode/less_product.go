package leetcode

func NumSubarrayProductLessThanK(nums []int, k int) int {
	num := 0
	product := 1

	for s, e := 0, 0; e < len(nums); e++ {
		product *= nums[e]

		for product >= k && s <= e {
			product /= nums[s]
			s++
		}

		num += e - s + 1
	}

	return num
}
