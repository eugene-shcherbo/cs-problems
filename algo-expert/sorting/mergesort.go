package sorting

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	left := MergeSort(array[:len(array)/2])
	right := MergeSort(array[len(array)/2:])

	return Merge(left, right)
}

func Merge(left, right []int) []int {
	res := make([]int, 0)

	for l, r := 0, 0; l < len(left) || r < len(right); {
		if l >= len(left) {
			res = append(res, right[r])
			r++
		} else if r >= len(right) {
			res = append(res, left[l])
			l++
		} else if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}

	return res
}
