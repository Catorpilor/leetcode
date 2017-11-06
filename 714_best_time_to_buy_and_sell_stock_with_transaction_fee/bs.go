package bs

import (
	"github.com/catorpilor/leetcode/utils"
)

func MaxProfit(prices []int, fee int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	buy, sell := make([]int, n), make([]int, n)
	buy[0] = 0 - prices[0]
	for i := 1; i < n; i++ {
		buy[i] = utils.Max(buy[i-1], sell[i-1]-prices[i])
		sell[i] = utils.Max(sell[i-1], buy[i-1]+prices[i]-fee)
	}
	return sell[n-1]
}

func MaxProfit2(prices []int, fee int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	buy, sell, prev_buy, prev_sell := 0, 0, -prices[0], 0
	for i := 1; i < n; i++ {
		buy = utils.Max(prev_buy, prev_sell-prices[i])
		sell = utils.Max(prev_sell, prev_buy+prices[i]-fee)
		prev_buy, prev_sell = buy, sell
	}
	return sell
}
