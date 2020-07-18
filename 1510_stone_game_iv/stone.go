package stone

func winnerSquareGame(n int) bool {
	return useDP(n)
}

// useDP time complexity O(N^3/2), space complexity O(N)
func useDP(n int) bool {
	// dp[i] means the result for n = i.
	// if there is any k that dp[i - k * k] == false,
	// it means we can make the other lose the game,
	// so we can win the game  dp[i] = true.
	dp := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for k := 1; k*k <= i; k++ {
			if !dp[i-k*k] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
