package graphs

func ValidPathDfs(n int, edges [][]int, source int, destination int) bool {
	graph := buildGraph(edges)
	visited := make([]bool, n)

	return validPathImpl(graph, visited, source, destination)
}

func validPathImpl(graph map[int][]int, visited []bool, src, dest int) bool {
	if src == dest {
		return true
	}

	if visited[src] {
		return false
	}

	visited[src] = true

	for _, e := range graph[src] {
		if validPathImpl(graph, visited, e, dest) {
			return true
		}
	}

	return false
}

func ValidPathBfs(n int, edges [][]int, source int, destination int) bool {
	graph := buildGraph(edges)
	visited := make([]bool, n)

	nextVertexes := make([]int, 0)
	nextVertexes = append(nextVertexes, source)

	for len(nextVertexes) > 0 {
		vertex := nextVertexes[0]
		nextVertexes = nextVertexes[1:]

		if vertex == destination {
			return true
		}

		for _, neigh := range graph[vertex] {
			if !visited[neigh] {
				visited[neigh] = true
				nextVertexes = append(nextVertexes, neigh)
			}
		}
	}

	return false
}

func buildGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)

	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	return graph
}
