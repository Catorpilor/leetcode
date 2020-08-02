package maxp

import (
	"fmt"
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func maxProb(n int, edges [][]int, succProb []float64, start, end int) float64 {
	return useDfs(n, edges, succProb, start, end)
}

func useDfs(n int, edges [][]int, succProb []float64, start, end int) float64 {
	// build adj list
	adj := make(map[int][]int, n)
	ne := len(edges)
	if ne < 1 {
		return float64(0)
	}
	cost := make(map[string]float64, ne)
	for i, edge := range edges {
		s, e := edge[0], edge[1]
		key := fmt.Sprintf("%d-%d", s, e)
		key1 := fmt.Sprintf("%d-%d", e, s)
		adj[s] = append(adj[s], e)
		adj[e] = append(adj[e], s)
		cost[key] = succProb[i]
		cost[key1] = succProb[i]
	}
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("%d-%d", i, i)
		cost[key] = float64(1)
	}
	visited := make([]bool, n)
	ans := float64(0)
	dfs(&ans, adj, cost, visited, start, start, end, float64(1))
	return ans
}

func dfs(ret *float64, adj map[int][]int, cost map[string]float64, visited []bool, prev, cur, end int, curP float64) {
	// fmt.Printf("edge(%d-%d), curP:%f\n", prev, cur, curP)
	visited[cur] = true
	if cur == end {
		*ret = math.Max(*ret, curP)
		// fmt.Printf("reached at end, curP:%f, ret:%f\n", curP, *ret)
	} else {
		for _, node := range adj[cur] {
			if !visited[node] {
				dfs(ret, adj, cost, visited, cur, node, end, curP*cost[fmt.Sprintf("%d-%d", cur, node)])
			}
		}
	}
	visited[cur] = false
	preKey := fmt.Sprintf("%d-%d", prev, cur)
	// fmt.Printf("preKey:%s, cost[preKey]=%f, curP:%f\n",preKey, cost[preKey], curP)
	curP /= cost[preKey]
}

// useBellmanFord time complexity O(|V|*|E|), space compelxity O(N*len(edges))
func useBellmanFord(n int, edges [][]int, succProb []float64, start, end int) float64 {
	// adj
	g := make(map[int][][]int, n)
	for i, edge := range edges {
		start, end := edge[0], edge[1]
		g[start] = append(g[start], []int{end, i})
		g[end] = append(g[end], []int{start, i})
	}
	costs := make([]float64, n)
	costs[start] = 1.0
	q := utils.NewQueue()
	q.Enroll(start)
	for !q.IsEmpty() {
		cur := q.Pull().(int)
		for _, pair := range g[cur] {
			neighbor, idx := pair[0], pair[1]
			if costs[cur]*succProb[idx] > costs[neighbor] {
				costs[neighbor] = costs[cur] * succProb[idx]
				q.Enroll(neighbor)
			}
		}
	}
	return costs[end]
}
