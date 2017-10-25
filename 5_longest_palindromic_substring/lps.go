package lps

func Lps(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	start, maxLen := 0, 1
	// for lenght 1
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	// for length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			start, maxLen = i, 2
			dp[i][i+1] = true
		}
	}
	// for length >= 3
	for cl := 3; cl <= n; cl++ {
		for i := 0; i < n-cl+1; i++ {
			j := i + cl - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start = i
				maxLen = cl
			}
		}
	}
	return s[start : start+maxLen]
}
