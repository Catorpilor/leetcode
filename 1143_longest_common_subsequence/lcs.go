package lcs

import "github.com/catorpilor/LeetCode/utils"

func longestCommonSeq(text1, text2 string) int {
	m, n := len(text1), len(text2)
	if text1 == text2 {
		return m
	}
	dp := make([]int, n+1)
	var upleft int
	for i := 1; i <= m; i++ {
		upleft = dp[0]
		for j := 1; j <= n; j++ {
			prev, cur := dp[j-1], dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = upleft + 1
			} else {
				dp[j] = utils.Max(prev, cur)
			}
			upleft = cur
		}
	}
	return dp[n]
}
