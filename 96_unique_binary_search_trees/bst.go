package bst

func numOfBst(n int) int {
	return useDP(n)
}

// useDP time complexity O(N^2), space complexity O(N)
func useDP(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[0] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
