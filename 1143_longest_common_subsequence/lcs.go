package lcs

import "github.com/catorpilor/leetcode/utils"

func longestCommonSeq(text1, text2 string) int {
	return useDPLessSpace(text1, text2)
}

// useDP time complexity O(MN), space complexity O(MN)
func useDP(text1, text2 string) int {
	n1, n2 := len(text1), len(text2)
	if n1 == 0 || n2 == 0 {
		return 0
	}
	if n1 > n2 {
		text1, text2 = text2, text1
		n1, n2 = n2, n1
	}
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = utils.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n1][n2]
}

// useDPLessSpace time complexity is O(mn)
// space compelxity is O(n)
func useDPLessSpace(text1, text2 string) int {
	m, n := len(text1), len(text2)
	if m == 0 || n == 0 {
		return 0
	}
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
