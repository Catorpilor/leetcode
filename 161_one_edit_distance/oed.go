package oed

import "github.com/catorpilor/leetcode/utils"

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

// onePass sovle using iteration
// time complexity O(m)
// space complexity O(1)
func onePass(s, t string) bool {
	m, n := len(s), len(t)
	if m > n {
		// always make len(s) <= len(t)
		return onePass(t, s)
	}
	if n-m > 1 {
		// no way
		return false
	}
	for i := 0; i < m; i++ {
		if s[i] != t[i] {
			if m == n {
				// len(s) == len(t) && s[[i] != t[i]
				//  this can only be the difference
				// so we have to check the remaining parts
				return s[i+1:] == t[i+1:]
			}
			// len(s) != len(t) && s[i] != t[i]
			// skip t[i] here and check s[i:] with t[i+1:]
			return s[i:] == t[i+1:]
		}
	}
	// if no difference found
	// just make sure m+1 == n
	return m+1 == n
}
