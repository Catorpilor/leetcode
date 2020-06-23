package dungeon

import (
	"math"

	"github.com/catorpilor/LeetCode/utils"
)

func calculateMinHP(grid [][]int) int {
	return useDP(grid)
}

// useDP time complexity O(MN), space complexity O(MN)
func useDP(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	// to reach (m,n-1), (m-1, n) we only need 1 hp.
	dp[m][n-1] = 1
	dp[m-1][n] = 1
	// Use dp[i][j] to store the min hp needed at position (i, j)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			need := utils.Min(dp[i+1][j], dp[i][j+1]) - grid[i][j]
			if need <= 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = need
			}
		}
	}
	return dp[0][0]
}
