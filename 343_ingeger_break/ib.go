package ib

import "github.com/catorpilor/leetcode/utils"

func IB(n int) int {
	if n < 2 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = utils.Max(dp[i], utils.Max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

// IB2 returns the max product for n break into n1--nk with sum as n
// the math behind it
// suppose for n we got an x that got the max product.
// y = x ^ (n/x)
// then y' = x^(n/x) * n/x^2 * (1-lnx)
// so if x = e makes it a peak,when x > e it became negative
// so for integers we have to choose 2 < e < 3 that makes the biggest product.
// but prefer more 3's cause 6 = 2*2*2 < 3*3
// but for n = 4 we got 2*2 > 3*1
func IB2(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	ret := 1
	for n > 4 {
		ret *= 3
		n -= 3
	}
	ret *= n
	return ret
}
