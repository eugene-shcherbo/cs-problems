package matrix

import "github.com/eugene-shcherbo/cs-problems/utils"

func MaxAreaOfIsland(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	maxArea := 0
	for x, r := range grid {
		for y, c := range r {
			if c == 1 && !visited[x][y] {
				maxArea = utils.Max(maxArea, calcIslandArea(grid, visited, x, y))
			}
		}
	}

	return maxArea
}

func calcIslandArea(grid [][]int, visited [][]bool, x, y int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
		return 0
	}

	if visited[x][y] {
		return 0
	}

	visited[x][y] = true

	area := calcIslandArea(grid, visited, x+1, y)
	area += calcIslandArea(grid, visited, x-1, y)
	area += calcIslandArea(grid, visited, x, y+1)
	area += calcIslandArea(grid, visited, x, y-1)

	return area + 1
}
