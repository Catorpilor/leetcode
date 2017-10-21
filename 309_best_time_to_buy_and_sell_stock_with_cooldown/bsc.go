package bsc

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

// MaxProfit returns the maxprofit for prices
func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	// buy[i] represents the max profit before day i which end with buying at i
	// sell[i] represents the max profit before day i which end with selling at i
	// rest[i] represents the max profit before day i which end with resting at i.
	// so we have
	// buy[i] = max(buy[i-1],rest[i-1]-prices[i])
	// sell[i] = max(sell[i-1],buy[i-1]+price[i])
	// rest[i] = max(sell[i-1],buy[i-1],rest[i-1])
	// for rest[i] we can ensure that it equals to sell[i-1]
	// so the aformentioned three equations can be rewirte to 2
	// buy[i] = max(buy[i-1],sell[i-2]-prices[i])
	// sell[i] = max(sell[i-1],buy[i-1]+prices[i])
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

func MaxProfit2(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	// since the aformentioned two euqations have to do with only three variables
	// so we can reduce the O(n) space to O(1)
	var prev_buy, prev_sell, buy, sell int
	buy = math.MinInt32
	for _, p := range prices {
		prev_buy = buy
		buy = utils.Max(prev_buy, prev_sell-p)
		prev_sell = sell
		sell = utils.Max(prev_sell, prev_buy+p)
	}
	return sell
}
