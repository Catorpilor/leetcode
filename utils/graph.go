package utils

// Graph represent a graph
type Graph struct {
	n int
	g map[int][]int
}

// ConstructAdjacencyList create an adjacency list based on the connections array
func ConstructAdjacencyList(n int, conns [][]int, directed bool) *Graph {
	ret := &Graph{g: make(map[int][]int, n), n: n}
	for _, c := range conns {
		u, v := c[0], c[1]
		ret.g[u] = append(ret.g[u], v)
		if !directed {
			ret.g[v] = append(ret.g[v], u)
		}
	}
	return ret
}

// NumofSccs returns the number of Strong Connected Componnets in the graph
func (gp *Graph) NumofSccs() int {
	return len(useTarjan(gp.g, gp.n))
}

func useTarjan(g map[int][]int, n int) [][]int {
	var res [][]int
	low, disc := make([]int, n), make([]int, n)
	for i := range disc {
		// use disc[i] = -1 to distinguish visited node
		disc[i] = -1
	}
	var idx int
	for i := range disc {
		if disc[i] == -1 {
			helper(g, low, disc, i, i, &idx)
		}
	}
	group := make(map[int][]int, n)
	for i := range low {
		group[low[i]] = append(group[low[i]], i)
	}
	// fmt.Printf("low: %v and group: %v\n", low, group)
	for k := range group {
		res = append(res, group[k])
	}
	return res
}

func helper(g map[int][]int, low, disc []int, cur, prev int, idx *int) {
	low[cur], disc[cur] = (*idx), (*idx)
	(*idx)++
	for _, j := range g[cur] {
		if j == prev {
			continue
		}
		if disc[j] == -1 {
			helper(g, low, disc, j, cur, idx)
			low[cur] = Min(low[cur], low[j])
		} else {
			low[cur] = Min(low[cur], disc[j])
		}
	}
}
