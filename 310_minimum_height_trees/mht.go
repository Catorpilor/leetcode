package mht

func findMinHeightTrees(n int, edges [][]int) []int {
	return nil
}

func cutTheLeaves(n int, edges [][]int) []int {
	ans := make([]int, 0, n)
	if n < 1 {
		return ans
	}
	if n == 1 {
		ans = append(ans, 0)
		return ans
	}
	adj := make(map[int][]int, n)
	for _, edge := range edges {
		s, e := edge[0], edge[1]
		adj[s] = append(adj[s], e)
		adj[e] = append(adj[e], s)
	}
	leaves := make([]int, 0, n)
	for k, v := range adj {
		if len(v) == 1 {
			leaves = append(leaves, k)
		}
	}
	for n > 2 {
		n -= len(leaves)
		// cut the leaves
		nl := make([]int, 0, n)
		for _, l := range leaves {
			other := adj[l][0]
			delete(adj, l)
			var i int
			for i = range adj[other] {
				if adj[other][i] == l {
					break
				}
			}
			adj[other] = append(adj[other][:i], adj[other][i+1:]...)
			if len(adj[other]) == 1 {
				nl = append(nl, other)
			}
		}
		leaves = nl
	}
	return leaves
}
