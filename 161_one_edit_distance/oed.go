package oed

import "github.com/catorpilor/LeetCode/utils"

// isOneEditDistance using dp to solve this problem
// time complexity O(mn)
// space complexity O(n)
func isOneEditDistance(s, t string) bool {
	m, n := len(s), len(t)
	if s == t {
		return false
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
	}
	var upleft int
	for i := 1; i <= m; i++ {
		upleft = dp[0]
		dp[0] = i
		for j := 1; j <= n; j++ {
			pre, cur := dp[j-1], dp[j]
			if s[i-1] == t[j-1] {
				dp[j] = upleft
			} else {
				dp[j] = utils.Min(upleft, utils.Min(pre, cur)) + 1
			}
			upleft = cur
		}
	}
	return dp[n] == 1
}
