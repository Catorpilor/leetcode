package keys

import (
	"github.com/catorpilor/leetcode/utils"
)

func MaxA(n int) int {
	if n < 7 {
		return n
	}
	dp := make([]int, n+1)
	for i := 0; i < 7; i++ {
		dp[i] = i
	}
	var maxv int
	for i := 7; i <= n; i++ {
		maxv = i
		for j, k := 6, 3; j <= i; j, k = j+1, k+1 {
			// k means first k 'A's
			// 2*dp[k] means k 'A's + 1 copy past of dp[k]
			// that's why j starts at 6 ( 3 'A', ctrl+a, ctrl+c, ctrl+v)
			// (i-j)*dp[k] means how many ctrlv will be added
			temp := 2*dp[k] + (i-j)*dp[k]
			if temp > maxv {
				maxv = temp
			}
		}
		dp[i] = maxv
	}
	return dp[n]
}

func MaxA2(n int) int {
	if n <= 6 {
		return n
	}
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = i
		for j := 3; j < i; j++ {
			// j represent ctrl+a, ctrl+c, ctrl+v, ctrl+v ....
			// for example j = 3 we have just one ctrl+v, so it equals to 2 times (j-1) of dp[i-3]
			dp[i] = utils.Max(dp[i], dp[i-j]*(j-1))
		}
	}
	return dp[n]
}

func MaxA3(n int) int {
	if n <= 6 {
		return n
	}
	dp := make([]int, n+1)
	for i := 0; i <= 6; i++ {
		dp[i] = i
	}
	for i := 7; i <= n; i++ {
		// no need to recaculate all j range [3,i)
		// j = 4, 5 are always the largest number in the series.
		dp[i] = utils.Max(dp[i-4]*3, dp[i-5]*4)
	}
	return dp[n]
}
