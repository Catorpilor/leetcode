package cc

import (
	"github.com/catorpilor/leetcode/utils"
)

func criticalConns(n int, conn [][]int) [][]int {
	return useTarjan(n, conn)
}

// useTarjan time complexity O(N+E), space complexity O(N+E)
func useTarjan(n int, conn [][]int) [][]int {
	var res [][]int
	low, disc := make([]int, n), make([]int, n)
	for i := range disc {
		// use disc[i] to track visited status
		disc[i] = -1
	}
	// build adjancency list
	g := make(map[int][]int, n)
	for _, c := range conn {
		h, t := c[0], c[1]
		g[h] = append(g[h], t)
		g[t] = append(g[t], h)
	}
	var idx int
	for i := range disc {
		if disc[i] == -1 {
			helper(&res, g, low, disc, i, i, &idx)
		}
	}
	return res
}

func helper(res *[][]int, g map[int][]int, low, disc []int, cur, pre int, idx *int) {
	low[cur], disc[cur] = (*idx), (*idx)
	(*idx)++
	// fmt.Printf("low: %v, disc: %v, cur: %d, pre: %d, idx: %d\n", low, disc, cur, pre, *idx)
	for _, j := range g[cur] {
		// AHA!
		// this is the key point, cause Tarjan algorithm is find the strongest connected components in
		// directed graph, we skip the u->v and v->u case
		if j == pre {
			continue
		}
		if disc[j] == -1 {
			helper(res, g, low, disc, j, cur, idx)
			low[cur] = utils.Min(low[cur], low[j])
			if low[j] > disc[cur] {
				// j do not belong to the SCC that includes cur
				*res = append(*res, []int{cur, j})
			}
		} else {
			// j is already in the SCC, update cur's lowindex if necessary
			low[cur] = utils.Min(low[cur], disc[j])
		}
	}
}
