package graph

func isBipartite(graph [][]int) bool {
	return useDFS(graph)
}

// useDFS time complexity O(N*|E|), space complexity O(N)
func useDFS(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)
	for i := range graph {
		if colors[i] == 0 && !helper(graph, colors, 1, i) {
			return false
		}
	}
	return true
}

func helper(graph [][]int, colors []int, color, node int) bool {
	if colors[node] != 0 {
		return colors[node] == color
	}
	// mark node with color
	colors[node] = color
	// traverse node's adjances
	for _, adj := range graph[node] {
		if !helper(graph, colors, -color, adj) {
			return false
		}
	}
	return true
}
