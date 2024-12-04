package graphs

func FindSmallestSetOfVertices(n int, edges [][]int) []int {
	inDegrees := make([]int, n)

	for _, edge := range edges {
		inDegrees[edge[1]]++
	}

	var result []int
	for n, inDegree := range inDegrees {
		if inDegree == 0 {
			result = append(result, n)
		}
	}

	return result
}
