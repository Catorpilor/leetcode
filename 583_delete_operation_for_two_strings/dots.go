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

// diff **wrong** solution
// for example word1="eat" word2="are"
// diff returns 2 but the result is 4
// it's a common lcs problem.
func diff(word1, word2 string) int {
	if word1 == word2 {
		return 0
	}
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	hset := make(map[rune]int, m+n)
	for _, b := range word1 {
		hset[b]++
	}
	for _, b := range word2 {
		if _, exists := hset[b]; exists {
			hset[b]--
			if hset[b] == 0 {
				delete(hset, b)
			}
		} else {
			hset[b]++
		}
	}
	return len(hset)
}
