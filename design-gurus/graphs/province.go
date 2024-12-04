package graphs

func CountProvinces(isConnected [][]int) int {
	isVisited := make([]bool, len(isConnected))
	provinces := 0

	for r := 0; r < len(isConnected); r++ {
		if isVisited[r] {
			continue
		}

		for c := 0; c < len(isConnected[r]); c++ {
			if isConnected[r][c] == 1 {
				provinces++
				visitProvince(c, isConnected, isVisited)
				break
			}
		}
	}

	return provinces
}

func visitProvince(row int, isConnected [][]int, isVisited []bool) {
	if isVisited[row] {
		return
	}

	isVisited[row] = true

	for c := 0; c < len(isConnected[row]); c++ {
		if isConnected[row][c] == 1 {
			visitProvince(c, isConnected, isVisited)
		}
	}
}
