package arrays

func RunningSum(nums []int) []int {
	res := make([]int, len(nums))

	sum := 0
	for i, el := range nums {
		sum += el
		res[i] = sum
	}

	return res
}
