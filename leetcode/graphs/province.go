package graphs

func FindCircleNum(isConnected [][]int) int {
	provinces := 0
	visited := make([]bool, len(isConnected))

	for i := range isConnected {
		if !visited[i] {
			provinces++
			visitProvince(isConnected, visited, i)
		}
	}

	return provinces
}

func visitProvince(isConnected [][]int, visited []bool, city int) {
	if visited[city] {
		return
	}

	visited[city] = true
	for neigh, connection := range isConnected[city] {
		if connection == 1 {
			visitProvince(isConnected, visited, neigh)
		}
	}
}
