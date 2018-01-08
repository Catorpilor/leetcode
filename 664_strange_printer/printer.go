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

func StrangePrinter2(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	dp := [101][101]int{}
	// for single character set dp[i][i] to 1
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	var right int
	for l := 2; l <= n; l++ {
		for left := 0; left < n-l+1; left++ {
			right = left + l - 1
			dp[left][right] = l
			for k := left; k < right; k++ {
				steps := dp[left][k] + dp[k+1][right]
				if s[k] == s[right] {
					steps--
				}
				dp[left][right] = utils.Min(dp[left][right], steps)
			}
		}
	}
	return dp[0][n-1]
}

func StrangePrinter3(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	dp := [101][101]int{}
	// for single character set dp[i][i] to 1
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	var right int
	for l := 2; l <= n; l++ {
		for left := 0; left < n-l+1; left++ {
			right = left + l - 1
			dp[left][right] = dp[left][right-1] + 1
			if s[right] == s[right-1] {
				dp[left][right]--
			}
			for k := left + 1; k <= right; k++ {
				if s[k-1] == s[right] {
					dp[left][right] = utils.Min(dp[left][right], dp[left][k-1]+dp[k][right]-1)
				}
			}
		}
	}
	return dp[0][n-1]
}
