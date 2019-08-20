package maze

import (
	"fmt"
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func CherryPickUp(grid [][]int) int {
	n := len(grid)
	if n == 1 {
		return grid[0][0]
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if grid[i-1][j-1] != -1 {
				dp[i][j] = grid[i-1][j-1] + utils.Max(dp[i-1][j], dp[i][j-1])
			} else {
				dp[i][j] = math.MinInt32
			}
		}
	}
	var ret int
	fmt.Printf("from 0,0 to n-1,n-1 we got %d\n", dp[n][n])
	if !validDownsidePath(grid, dp, n, n) {
		return ret
	}
	ret += dp[n][n]
	fmt.Printf("now grid %v\n", grid)
	dp2 := make([][]int, n+1)
	for i := range dp2 {
		dp2[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] != -1 {
				dp2[i][j] = grid[i][j] + utils.Max(dp2[i+1][j], dp2[i][j+1])
			} else {
				dp2[i][j] = math.MinInt32
			}
		}
	}
	ret += dp2[0][0]

	return ret
}

func validDownsidePath(grid [][]int, dp [][]int, i, j int) bool {
	for i >= 1 && j >= 1 {
		if dp[i-1][j] == math.MinInt32 && dp[i][j-1] == math.MinInt32 {
			return false
		}
		grid[i-1][j-1] = 0
		if dp[i-1][j] > dp[i][j-1] {
			// down
			i = i - 1
		} else {
			// from right
			j = j - 1
		}
	}
	return true
}
