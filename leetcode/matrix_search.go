package leetcode

func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, 0, len(matrix), 0, len(matrix[0]), target)
}

func searchMatrix(matrix [][]int, srow, erow, scol, ecol int, target int) bool {
	if srow > erow || scol > ecol {
		return false
	}

	mRow := (srow + erow) / 2
	mCol := (scol + ecol) / 2

	if matrix[mRow][mCol] == target {
		return true
	} else if matrix[mRow][mCol] > target {
		return searchMatrix(matrix, srow, mRow-1, scol, ecol, target) || searchMatrix(matrix, mRow, erow, scol, mCol-1, target)
	} else {
		return searchMatrix(matrix, mRow+1, erow, scol, ecol, target) || searchMatrix(matrix, srow, mRow, mCol+1, ecol, target)
	}
}
