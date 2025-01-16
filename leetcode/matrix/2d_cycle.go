package matrix

func ContainsCycle(grid [][]byte) bool {
	visited := make([][]bool, len(grid))
	for x, r := range grid {
		visited[x] = make([]bool, len(r))
	}

	for x, r := range grid {
		for y := range r {
			if !visited[x][y] && formsCycle(grid, visited, x, y, x, y) {
				return true
			}
		}
	}

	return false
}

func formsCycle(grid [][]byte, visited [][]bool, x, y int, prevX, prevY int) bool {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return false
	}

	if grid[x][y] != grid[prevX][prevY] {
		return false
	}

	if visited[x][y] {
		return true
	}

	visited[x][y] = true

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, d := range dirs {
		newX := x + d[0]
		newY := y + d[1]

		if (newX != prevX || newY != prevY) && formsCycle(grid, visited, newX, newY, x, y) {
			return true
		}
	}

	return false
}
