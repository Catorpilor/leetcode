package dots

import "github.com/catorpilor/LeetCode/utils"

// minDistance using dp to solve the problem.
// time complexity is O(mn)
// space complexity is o(n)
func minDistance(word1, word2 string) int {
	if word1 == word2 {
		return 0
	}
	m, n := len(word1), len(word2)
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
			if word1[i-1] == word2[j-1] {
				dp[j] = upleft
			} else {
				dp[j] = utils.Min(pre, cur) + 1
			}
			upleft = cur
		}
	}
	return dp[n]
}
