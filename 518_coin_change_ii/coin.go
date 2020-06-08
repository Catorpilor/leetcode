package coin

func change(amount int, coins []int) int {
	return useDP(amount, coins)
}

// useDP time complexity O(MN), space complexity O(MN)
func useDP(amount int, coins []int) int {
	n := len(coins)
	if n == 0 || (n == 1 && amount%coins[0] != 0) {
		return 0
	}
	dp := make([][]int, amount+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = 1 // always one way to choose whem amount = 0
	}
	for i := 1; i <= amount; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i][j-1] // exclude this coin
			if i-coins[j-1] >= 0 {
				dp[i][j] += dp[i-coins[j-1]][j] // include this coin
			}
		}
	}
	return dp[amount][n]
}
