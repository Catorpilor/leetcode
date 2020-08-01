package climb

func numOfClimbs(n int) int {
	return useDP(n)
}

// useDP time complexity O(N), space complexity O(N)
func useDP(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

// useLessSpaceDP time complexity O(N), space complexity O(1)
func useLessSpaceDP(n int) int {
	if n <= 2 {
		return n
	}
	minusOne, minusTwo := 1, 2
	var ans int
	for i := 3; i <= n; i++ {
		ans = minusOne + minusTwo
		minusOne = minusTwo
		minusTwo = ans
	}
	return ans
}
