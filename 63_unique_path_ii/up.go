package up

import "fmt"

func UniquePath(grid [][]int) int {
	if grid == nil {
		return 0
	}
	m, n := len(grid), len(grid[0])
	if grid[0][0] == 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	prev := dp[0]
	for j := 0; j < m; j++ {
		prev = 1
		if grid[j][0] == 1 || dp[0] == 0 {
			prev = 0
			dp[0] = 0
		}
		for i := 1; i < n; i++ {
			if grid[j][i] == 0 {
				dp[i] = prev + dp[i]
			} else {
				dp[i] = 0
			}
			prev = dp[i]
		}
		fmt.Println(dp)
	}
	return dp[n-1]
}

func UniquePath2(grid [][]int) int {
	if grid == nil {
		return 0
	}
	n := len(grid[0])
	dp := make([]int, n)
	dp[0] = 1
	for _, c := range grid {
		for j := 0; j < n; j++ {
			if c[j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}
