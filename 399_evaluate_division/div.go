package div

func calcEquation(eqs [][]string, vals []float64, queries [][]string) []float64 {
	// return useBfs(eqs, vals, queries)
	return useDFS(eqs, queries, vals)
}

type pair struct {
	node string
	val  float64
}

// useBfs time complexity O(N^2), space complexity O(N)
func useBfs(eqs [][]string, vals []float64, queires [][]string) []float64 {
	adj := make(map[string][]pair)
	cache := make(map[string]float64)
	for i, eq := range eqs {
		s, e, val := eq[0], eq[1], vals[i]
		// there is an edge from f->e with value val
		adj[s] = append(adj[s], pair{e, val})
		// there is an edge from e->ef with valus 1/val
		adj[e] = append(adj[e], pair{s, 1 / val})
	}
	ans := make([]float64, len(queires))
	for i, query := range queires {
		s, e := query[0], query[1]
		key := s + "/" + e
		if _, exists := cache[key]; !exists {
			helper(adj, cache, s, e)
		}
		ans[i] = cache[key]
	}
	return ans
}

func helper(adj map[string][]pair, cache map[string]float64, start, end string) {
	// start or end not in the graph
	if len(adj[start]) == 0 || len(adj[end]) == 0 {
		cache[start+"/"+end] = -1.0
		cache[end+"/"+start] = -1.0
		return
	}
	queue := make([]pair, 0, 20)
	queue = append(queue, pair{start, 1.0})
	visited := make(map[string]bool)
	for len(queue) > 0 {
		// simulate the pop commans of a queue
		top := queue[0]
		queue = queue[1:]
		node, val := top.node, top.val
		if node == end {
			cache[start+"/"+end] = val
			cache[end+"/"+start] = 1 / val
			return
		}
		for _, next := range adj[node] {
			if !visited[next.node] {
				visited[next.node] = true
				queue = append(queue, pair{next.node, next.val * val})
			}
		}
	}
	cache[start+"/"+end] = -1.0
	return
}

// useDFS time complexity O(KN) where k is the number of queries N is the number of nodes. Space complexity O(N)
func useDFS(eqs, queries [][]string, vals []float64) []float64 {
	adj := make(map[string][]pair)
	for i, eq := range eqs {
		s, e, val := eq[0], eq[1], vals[i]
		// there is an edge from f->e with value val
		adj[s] = append(adj[s], pair{e, val})
		// there is an edge from e->ef with valus 1/val
		adj[e] = append(adj[e], pair{s, 1 / val})
	}
	ans := make([]float64, len(queries))
	for i := range ans {
		ans[i] = -1.0
	}
	for i, query := range queries {
		calWeight(query[0], query[1], make(map[string]bool), adj, &ans, i, 1.0)
	}
	return ans
}

func calWeight(start, end string, visited map[string]bool, adj map[string][]pair, ans *[]float64, idx int, val float64) {
	if len(adj[start]) == 0 || len(adj[end]) == 0 {
		return
	}
	if start == end {
		(*ans)[idx] = val
		return
	}
	visited[start] = true
	for _, next := range adj[start] {
		if !visited[next.node] {
			calWeight(next.node, end, visited, adj, ans, idx, val*next.val)
		}
	}
}
