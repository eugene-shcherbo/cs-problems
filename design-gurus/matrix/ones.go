package matrix

func RowWithMaxOnes(matrix [][]int) (int, int) {
	if len(matrix) == 0 {
		return -1, -1
	}

	rowIdx := 0
	maxOnes := 0

	for i, row := range matrix {
		ones := 0
		for _, val := range row {
			if val == 1 {
				ones++
			}
		}

		if ones > maxOnes {
			maxOnes = ones
			rowIdx = i
		}
	}

	return rowIdx, maxOnes
}
