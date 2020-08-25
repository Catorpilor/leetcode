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
