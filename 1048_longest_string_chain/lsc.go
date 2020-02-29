package lsc

import "github.com/catorpilor/leetcode/utils"

func longestChain(words []string) int {
	return useGraph(words)
}

// useGraph time complexity O(N*L^2), space complexity O(N)
func useGraph(words []string) int {
	hset := make(map[string]int, len(words))
	g := make(map[int][]int, len(words))
	// O(N)
	for i := range words {
		hset[words[i]] = i
	}
	// O(N*L^2) build adjacency list
	for i, w := range words {
		for j := 0; j < len(w); j++ { // O(L)
			nw := w[:j] + w[j+1:] // O(L)
			if idx, exists := hset[nw]; exists {
				g[idx] = append(g[idx], i)
			}
		}
	}
	res := 0
	ans := make([]int, len(words))
	for i := range words {
		helper(g, i, ans)
		res = utils.Max(res, ans[i])
	}
	return res
}

// helper use memorization to calculate the longest chain from idx
func helper(g map[int][]int, idx int, ans []int) {
	if ans[idx] > 0 {
		return
	}
	ans[idx] = 1
	// loop through its connected pos
	for _, pos := range g[idx] {
		helper(g, pos, ans)
		ans[idx] = utils.Max(ans[idx], ans[pos]+1)
	}
}
