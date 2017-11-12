package keys

func MinStep(n int) int {
	if n <= 1 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 2; j < i; j++ {
			if i%j == 0 {
				dp[i] = j + dp[i/j]
				break
			}
		}
	}
	return dp[n]
}

func MinStep2(n int) int {
	if n <= 1 {
		return 0
	}
	var ret int
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			ret, n = ret+i, n/i
		}
	}
	return ret
}
