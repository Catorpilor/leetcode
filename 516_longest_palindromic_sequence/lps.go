package lps

import (
	"github.com/catorpilor/leetcode/utils"
)

func Lps(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	// for length = 2
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			dp[i-1][i] = 2
		} else {
			dp[i-1][i] = 1
		}
	}

	for l := 3; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				dp[i][j] = utils.Max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][n-1]
}
