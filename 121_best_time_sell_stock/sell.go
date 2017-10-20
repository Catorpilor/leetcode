package sell

import "github.com/catorpilor/leetcode/utils"

func Sell(prices []int) int {
	maxCur := 0
	maxSofar := 0
	for i := 1; i < len(prices); i++ {
		maxCur += prices[i] - prices[i-1]
		maxCur = utils.Max(0, maxCur)
		maxSofar = utils.Max(maxCur, maxSofar)
	}
	return maxSofar
}
