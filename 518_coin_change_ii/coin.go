package coin

func change(amount int, coins []int) int {
	return useDP(amount, coins)
}

// useBF time compleixty O(2^N), space complexity O(N)
func useBF(amount int, coins []int) int {
	return helper(amount, coins, 0)
}

func helper(amount int, coins []int, idx int) int {
	if amount < 0 {
		return 0
	}
	if amount == 0 {
		// find one solution
		return 1
	}
	if amount > 0 && idx >= len(coins) {
		// can not find a solution
		return 0
	}
	// include this coin + exclude this coin
	return helper(amount-coins[idx], coins, idx) + helper(amount, coins, idx+1)
}

// useDP time complexity O(MN), space complexity O(MN)
func useDP(amount int, coins []int) int {
	n := len(coins)
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
