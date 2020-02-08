package mad

import "github.com/catorpilor/leetcode/utils"

// minDeleteSum use dp to solve the problem
// time complexity o(mn)
// space complexity O(n)
func minDeleteSum(s1, s2 string) int {
	if s1 == s2 {
		return 0
	}
	m, n := len(s1), len(s2)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + int(s2[i-1])
	}
	var upleft int
	for i := 1; i <= m; i++ {
		upleft = dp[0]
		dp[0] += int(s1[i-1])
		for j := 1; j <= n; j++ {
			pre, cur := dp[j-1], dp[j]
			if s1[i-1] == s2[j-1] {
				dp[j] = upleft
			} else {
				dp[j] = utils.Min(pre+int(s2[j-1]), cur+int(s1[i-1]))
			}
			upleft = cur
		}
	}
	return dp[n]
}
