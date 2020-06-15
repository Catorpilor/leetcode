package kstops

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func findCheapest(n int, flights [][]int, src, dest, k int) int {
	return useBellmanFord(n, flights, src, dest, k)
}

func useBellmanFord(n int, flights [][]int, src, dest, k int) int {
	costs := make([]int, n)
	for i := range costs {
		costs[i] = math.MaxInt32
	}
	costs[src] = 0
	for s := 0; s <= k; s++ {
		local := make([]int, n)
		copy(local, costs)
		for _, p := range flights {
			if costs[p[0]] != math.MaxInt32 {
				local[p[1]] = utils.Min(local[p[1]], costs[p[0]]+p[2])
			}
		}
		costs = local
	}
	if costs[dest] == math.MaxInt32 {
		return -1
	}
	return costs[dest]
}
