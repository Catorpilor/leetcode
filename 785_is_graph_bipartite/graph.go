package graph

func isBipartite(graph [][]int) bool {
	// return useDFS(graph)
	return useBFS(graph)
}

// useDFS time complexity O(N), space complexity O(N)
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

// useBFS time complexity O(N), space complexity O(N)
func useBFS(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)
	q := make([]int, 0, n)
	// the graph may not be connected just traverse all the nodes.
	for i := range graph {
		if colors[i] != 0 {
			continue
		}
		colors[i] = 1
		q = append(q, i)
		for len(q) != 0 {
			top := q[0]
			q = q[1:]
			for _, adj := range graph[top] {
				if colors[adj] == 0 {
					colors[adj] = -colors[top]
					q = append(q, adj)
				} else if colors[adj] == colors[top] {
					return false
				}
			}
		}
	}
	return true
}
