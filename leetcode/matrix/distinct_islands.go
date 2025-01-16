package matrix

func NumDistinctIslands(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for x, r := range grid {
		visited[x] = make([]bool, len(r))
	}

	distinctIslands := make(map[string]bool)
	for x, r := range grid {
		for y, c := range r {
			if c == 1 && !visited[x][y] {
				path := visitIsland(grid, visited, x, y, "S")
				distinctIslands[path] = true
			}
		}
	}

	return len(distinctIslands)
}

func visitIsland(grid [][]int, visited [][]bool, x, y int, dir string) string {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return ""
	}

	if visited[x][y] {
		return ""
	}

	visited[x][y] = true

	path := dir
	path += visitIsland(grid, visited, x-1, y, "T")
	path += visitIsland(grid, visited, x+1, y, "B")
	path += visitIsland(grid, visited, x, y+1, "R")
	path += visitIsland(grid, visited, x, y-1, "L")
	path += "E"

	return path
}
