package path

func allPaths(graph [][]int) [][]int {
	return useDFS(graph)
}

func useDFS(graph [][]int) [][]int {
	n := len(graph)
	adj := make([][]int, n)
	for i := range graph {
		if len(graph[i]) > 0 {
			adj[i] = append(adj[i], graph[i]...)
		}
	}
	ans := make([][]int, 0, n)
	helper(&ans, adj, []int{0}, 0, n-1)
	return ans
}

func helper(ans *[][]int, adj [][]int, tmp []int, cur, target int) {
	if cur == target {
		local := make([]int, len(tmp))
		copy(local, tmp)
		*ans = append((*ans), local)
		return
	}
	for _, node := range adj[cur] {
		nt := len(tmp)
		tmp = append(tmp, node)
		helper(ans, adj, tmp, node, target)
		tmp = tmp[:nt]
	}
}
