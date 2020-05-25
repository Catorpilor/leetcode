package uncrossed

import "github.com/catorpilor/LeetCode/utils"

func maxLines(arr1, arr2 []int) int {
	return useDP(arr1, arr2)
}

// useDP time complexity O(MN), space complexity O(MN)
func useDP(arr1, arr2 []int) int {
	m, n := len(arr1), len(arr2)
	if m < 1 || n < 1 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if arr1[i-1] == arr2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = utils.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
