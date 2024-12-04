package arrays

func FindDiffArray(nums []int) []int {
	res := make([]int, len(nums))

	wholeSum := sumSlice(nums)
	runningSum := 0

	for i, num := range nums {
		wholeSum -= num
		res[i] = absInt(runningSum - wholeSum)
		runningSum += num
	}

	return res
}

func sumSlice(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
