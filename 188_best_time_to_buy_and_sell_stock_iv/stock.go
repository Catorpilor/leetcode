package stock

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func maxProfit(k int, prices []int) int {
	return useDP1(k, prices)
}

// useDP1 time compelxity O(N*K), space complexity O(N*K)
func useDP1(k int, prices []int) int {
	n := len(prices)
	if n < 1 || k == 0 {
		return 0
	}
	var ans int
	if k >= n>>1 {
		// meas you can make as much transactions as you want
		for i := 1; i < n; i++ {
			ans += utils.Max(prices[i]-prices[i-1], 0)
		}
		return ans
	}
	// k=2
	// dp[i][j] represents at most i transactions, the max profit ending at day j (means selling at day j)
	// dp[i][j] = max(dp[i][j-1], prices[j] - prices[m] + dp[i-1][m]) for m is [0.j-1]
	//   		= max(dp[i][j-1], prices[j] + max(dp[i-1][j] - prices[j]))
	// for example where i = 1, j = 2
	// m = 0 prices[j] - prices[0] + dp[i-1][0]
	// m = 1 prices[j] - prices[1] + dp[i-1][1]
	// suppose maxdif  for i = 1, j= 2 is on m = 1
	// for i = 1 j = 3
	// maxdif = utils.Max(maxdif, dp[i-1][j] - prices[j])

	// dp[i][0] = 0 means at day 0 no profit
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 1; i <= k; i++ {
		tmpMax := -prices[0]
		for j := 1; j < n; j++ {
			dp[i][j] = utils.Max(dp[i-1][j], tmpMax+prices[j])
			tmpMax = utils.Max(tmpMax, dp[i-1][j]-prices[j])
		}
	}
	return dp[k][n-1]
}

// useDP2 time complexity O(N*k), space compelxity O(K)
func useDP2(k int, prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	// simple case
	if k > n/2 {
		// means you can make any transactions you want
		var ans int
		for i := 1; i < n; i++ {
			ans += utils.Max(prices[i]-prices[i-1], 0)
		}
		return ans
	}
	buys := make([]int, k+1)
	sells := make([]int, k+1)
	for i := range buys {
		buys[i] = math.MinInt32
	}
	// buy1, buy2 := math.MinInt32, math.MinInt32
	// sel1, sel2 := 0, 0
	for _, p := range prices {
		// sel2 = utils.Max(sel2, buy2+p)
		// buy2 = utils.Max(buy2, sel1-p)
		// sel1 = utils.Max(sel1, buy1+p)
		// buy1 = utils.Max(buy1, -p)
		for j := k; j > 0; j-- {
			sells[j] = utils.Max(sells[j], buys[j]+p)
			buys[j] = utils.Max(buys[j], sells[j-1]-p)
		}
	}
	return sells[k]
}
