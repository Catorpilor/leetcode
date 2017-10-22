package up

func UniquePath(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	prev := dp[0]
	for j := 0; j < m; j++ {
		prev = 1
		for i := 1; i < n; i++ {
			dp[i] = prev + dp[i]
			prev = dp[i]
		}
	}
	return dp[n-1]

}
