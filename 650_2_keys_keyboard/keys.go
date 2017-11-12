package keys

func MinStep(n int) int {
	if n <= 1 {
		return 0
	}
	dp := make([]int, n+1)
	// dp[k] represents the **min steps** to get k 'A's
	// so to get min steps, the action must be ended with paste
	// suppose the last copy-all pressed when there are j 'A' s
	// to get k 'A's we have to (k-j)/j steps with paste
	// so the total steps to get k 'A's from j 'A's is 1 (copy-all) + (k-j)/j (paste)
	// so we have this formular dp[k] = j + dp[k/j]
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
