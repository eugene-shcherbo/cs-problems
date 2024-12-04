package matrix

func DiagSum(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	sum := 0
	mainIdx := 0
	secondIdx := len(matrix[0]) - 1

	for r := 0; r < len(matrix); r++ {
		sum += matrix[r][mainIdx]

		if mainIdx != secondIdx {
			sum += matrix[r][secondIdx]
		}

		mainIdx++
		secondIdx--
	}

	return sum
}
