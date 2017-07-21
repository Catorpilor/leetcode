package paint

func MinCost(costs [][]int) int {
	l := len(costs)
	if l == 0 {
		return 0
	}

	for i := 1; i < l; i++ {
		costs[i][0] += min(costs[i-1][1], costs[i-1][2]) // house i paint blue's costs
		costs[i][1] += min(costs[i-1][0], costs[i-1][2]) // house i paint blue's costs
		costs[i][2] += min(costs[i-1][1], costs[i-1][0]) // house i paint blue's costs
	}
	return min(costs[l-1][0], min(costs[l-1][1], costs[l-1][2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
