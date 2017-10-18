package cb

func CountingBits(n int) []int {
	if n < 0 {
		return nil
	}
	ret := make([]int, n+1)
	var i, cb int
	for ; i <= n; i++ {
		cb = 0
		num := i
		for num > 0 {
			num &= (num - 1)
			cb += 1
		}
		ret[i] = cb
	}
	return ret
}

func CountingBits2(num int) []int {
	if num < 0 {
		return nil
	}
	dp := make([]int, num+1)
	dp[0] = 0
	for i := 1; i <= num; i++ {
		dp[i] = dp[i>>1] + i&1
	}
	return dp
}

func CountingBits3(num int) []int {
	if num < 0 {
		return nil
	}
	ret := make([]int, num+1)
	for i := 1; i <= num; i++ {
		ret[i] = ret[i&(i-1)] + 1
	}
	return ret
}
