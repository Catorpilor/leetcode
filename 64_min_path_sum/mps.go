package mps

import "github.com/catorpilor/leetcode/utils"

func minPathSum(grid [][]int) int {
	return useDp(grid)
}

// useDp time complexity O(MN), space complexity O(MN)
func useDp(grid [][]int) int {
	m := len(grid)
	if m < 1 {
		return 0
	}
	n := len(grid[0])
	if n < 1 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		if j == 0 {
			dp[0][j] = grid[0][0]
		} else {
			dp[0][j] = dp[0][j-1] + grid[0][j]
		}
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i-1][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = utils.Min(dp[i-1][j]+grid[i-1][j], dp[i][j-1]+grid[i][j-1])
		}
	}
	return dp[m-1][n-1]
}
