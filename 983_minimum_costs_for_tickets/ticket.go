package ticket

import "github.com/catorpilor/leetcode/utils"

func minCostTickets(days []int, costs []int) int {
	return useDP(days, costs)
}

// useDP time complexity O(N), space complexity O(MAXDAY)
func useDP(days []int, costs []int) int {
	n := len(days)
	lastDay := days[n-1]
	set := make([]bool, lastDay+1)
	for _, d := range days {
		set[d] = true
	}
	dp := make([]int, lastDay+1)
	for i := 1; i <= lastDay; i++ {
		if !set[i] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1] + costs[0]
			j := 0
			if i >= 7 {
				j = i - 7
			}
			dp[i] = utils.Min(dp[i], dp[j]+costs[1])
			j = 0
			if i >= 30 {
				j = i - 30
			}
			dp[i] = utils.Min(dp[i], dp[j]+costs[2])
		}
	}
	return dp[lastDay]
}

// useQueue time complexity O(N), space complexity O(1)
func useQueue(days []int, costs []int) int {
	type pair struct {
		day  int
		cost int
	}
	lastSeven := make([]pair, 0, 7)
	lastMonth := make([]pair, 0, 30)
	var cost int
	for _, d := range days {
		for len(lastSeven) != 0 && lastSeven[0].day+7 <= d {
			lastSeven = lastSeven[1:]
		}
		for len(lastMonth) != 0 && lastMonth[0].day+30 <= d {
			lastMonth = lastMonth[1:]
		}
		lastSeven = append(lastSeven, pair{day: d, cost: cost + costs[1]})
		lastMonth = append(lastMonth, pair{day: d, cost: cost + costs[2]})
		cost = utils.Min(cost+costs[0], utils.Min(lastMonth[0].cost, lastSeven[0].cost))
	}
	return cost
}
