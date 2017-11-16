package regular

func IsMatch(s, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(p) == 1 {
		if len(s) != 1 {
			return false
		} else {
			return s[0] == p[0] || p[0] == '.'
		}
	}
	if p[1] == '*' {
		return IsMatch(s, p[2:]) || len(s) != 0 && (s[0] == p[0] || p[0] == '.') && IsMatch(s[1:], p)
	} else {
		return len(s) != 0 && (s[0] == p[0] || p[0] == '.') && IsMatch(s[1:], p[1:])
	}
}

func IsMatchDP(s, p string) bool {
	m, n := len(s), len(p)
	if m == 0 {
		return n == 0
	}
	// dp[i][j] represents s[0..i] and p[0..j] is match or not
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := range p {
		if i > 0 && p[i] == '*' && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[j] != '*' {
				dp[i+1][j+1] = dp[i][j] && (p[j] == '.' || p[j] == s[i])
			} else {
				// * case
				if p[j-1] != s[i] && p[j-1] != '.' {
					// repeat 0 time
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					// repeat al least 1 time
					dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j] || dp[i+1][j-1]
				}
				// dp[i+1][j+1] = dp[i+1][j-1] && j > 0 || dp[i+1][j] || dp[i][j+1] && j > 0 && (p[j-1] == '.' || s[i] == p[j-1])
			}
		}
	}
	return dp[m][n]

}
