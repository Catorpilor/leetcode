package bsc

import "github.com/catorpilor/leetcode/utils"

// MaxProfit returns the maxprofit for prices
func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	buy, sell := make([]int, len(prices)), make([]int, len(prices))
	buy[0] = -prices[0]
	buy[1] = utils.Max(buy[0], -prices[1])
	sell[1] = utils.Max(0, prices[1]-prices[0])
	for i := 2; i < len(prices); i++ {
		buy[i] = utils.Max(buy[i-1], sell[i-2]-prices[i])
		sell[i] = utils.Max(sell[i-1], buy[i-1]+prices[i])
	}
	return sell[len(prices)-1]
}
