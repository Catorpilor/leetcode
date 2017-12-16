package stock

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func Stock(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
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
	dp := make([][]int, 3)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 1; i <= 2; i++ {
		tempMax := dp[i-1][0] - prices[0]
		for j := 1; j < n; j++ {
			dp[i][j] = utils.Max(dp[i][j-1], prices[j]+tempMax)
			tempMax = utils.Max(tempMax, dp[i-1][j]-prices[j])
		}
	}
	return dp[2][n-1]
}

func Stock2(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	buy1, buy2 := math.MinInt32, math.MinInt32
	sel1, sel2 := 0, 0
	for _, p := range prices {
		sel2 = utils.Max(sel2, buy2+p)
		buy2 = utils.Max(buy2, sel1-p)
		sel1 = utils.Max(sel1, buy1+p)
		buy1 = utils.Max(buy1, -p)
	}
	return sel2
}
