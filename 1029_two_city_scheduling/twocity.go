package twocity

import "sort"

func minCosts(costs [][]int) int {
	return useGreedy(costs)
}

func useGreedy(costs [][]int) int {
	n := len(costs)
	var allInFirst int
	diffs := make([]int, n)
	for i, c := range costs {
		allInFirst += c[0]
		diffs[i] = c[1] - c[0]
	}
	sort.Slice(diffs, func(i, j int) bool {
		return diffs[i] < diffs[j]
	})
	for i := 0; i < n/2; i++ {
		allInFirst += diffs[i]
	}
	return allInFirst
}
