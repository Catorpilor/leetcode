package dot

import "github.com/catorpilor/leetcode/utils"

func maxDotProduct(num1, num2 []int) int {
	return useDP(num1, num2)
}

// useDP time complexity O(MN), space complexity O(MN)
func useDP(num1, num2 []int) int {
	m, n := len(num1), len(num2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = num1[i] * num2[j]
			if i > 0 && j > 0 {
				dp[i][j] += utils.Max(dp[i-1][j-1], 0)
			}
			if i > 0 {
				dp[i][j] = utils.Max(dp[i][j], dp[i-1][j])
			}
			if j > 0 {
				dp[i][j] = utils.Max(dp[i][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}
