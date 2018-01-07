package printer

import "github.com/catorpilor/leetcode/utils"

func StrangePrinter(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	// dp[i][j] represents the minimum steps to print s[i] to s[j]
	// dp[i][j] = dp[i][j-1] + 1
	// or dp[i][j] = min(dp[i][j], for every t in [i,j) if s[t] == s[j] dp[i][t]+dp[t+1][j-1]
	dp := [100][100]int{}
	return dfs(s, &dp, 0, n-1)

}

func dfs(s string, ans *[100][100]int, l, r int) int {
	if l > r {
		return 0
	}
	if (*ans)[l][r] != 0 {
		return (*ans)[l][r]
	}
	(*ans)[l][r] = dfs(s, ans, l, r-1) + 1
	for i := l; i < r; i++ {
		if s[i] == s[r] {
			(*ans)[l][r] = utils.Min((*ans)[l][r], dfs(s, ans, l, i)+dfs(s, ans, i+1, r-1))
		}
	}
	return (*ans)[l][r]
}
