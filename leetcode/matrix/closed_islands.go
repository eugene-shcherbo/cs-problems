package matrix

func ClosedIsland(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i, r := range grid {
		visited[i] = make([]bool, len(r))
	}

	closedIslands := 0
	for x, r := range grid {
		for y, c := range r {
			if c == 0 && !visited[x][y] && isClosedIsland(grid, visited, x, y) {
				closedIslands++
			}
		}
	}

	return closedIslands
}

func isClosedIsland(grid [][]int, visited [][]bool, x, y int) bool {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return false
	}

	if visited[x][y] || grid[x][y] == 1 {
		return true
	}

	visited[x][y] = true

	bot := isClosedIsland(grid, visited, x+1, y)
	top := isClosedIsland(grid, visited, x-1, y)
	right := isClosedIsland(grid, visited, x, y+1)
	left := isClosedIsland(grid, visited, x, y-1)

	return left && right && top && bot
}
