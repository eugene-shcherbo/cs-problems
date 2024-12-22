package leetcode

func MySqrt(x int) int {
	res := 0
	left := 1
	right := x

	for left <= right {
		mid := (left + right) / 2

		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}
