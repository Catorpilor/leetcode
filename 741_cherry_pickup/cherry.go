package maze

import (
	"fmt"
	"math"

	"github.com/catorpilor/leetcode/utils"
)

// Wrong Solution
// Greedy DP not working
// this is a two-leg dp, one leg greedy did not mean the global optimization
// for example
/*
  grid = [[1,1,1,0,1],
        [0,0,0,0,0],
        [0,0,0,0,0],
        [0,0,0,0,0],
        [1,0,1,1,1]].
   greedy solution: from (0,0) -> (n-1, n-1) is 6, from (n-1,n-1) -> (0,0) is 1 so the total is 7.
   The expected solution is follow the edges,
   (0,0) -> (n-1, n-1) is 5
   (n-1,n-1) -> (0,0) is 3, so the total is 8
*/
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

// This solution based on the discussion
// https://leetcode.com/problems/cherry-pickup/discuss/109903/Step-by-step-guidance-of-the-O(N3)-time-and-O(N2)-space-solution

func CherryPickUpNew(grid [][]int) int {
	N := len(grid)
	M := (N << 1) - 1
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, N)
	}
	dp[0][0] = grid[0][0]
	for n := 1; n < M; n++ {
		for i := N - 1; i >= 0; i-- {
			for p := N - 1; p >= 0; p-- {
				j, q := n-i, n-p
				if j < 0 || j >= N || q < 0 || q >= N || grid[i][j] < 0 || grid[p][q] < 0 {
					dp[i][p] = -1
					continue
				}

				if i > 0 {
					dp[i][p] = utils.Max(dp[i][p], dp[i-1][p])
				}
				if p > 0 {
					dp[i][p] = utils.Max(dp[i][p], dp[i][p-1])
				}
				if i > 0 && p > 0 {
					dp[i][p] = utils.Max(dp[i][p], dp[i-1][p-1])
				}
				if dp[i][p] >= 0 {
					dp[i][p] += grid[i][j]
					if i != p {
						dp[i][p] += grid[p][q]
					}
				}
			}
		}
	}
	return utils.Max(dp[N-1][N-1], 0)
}
