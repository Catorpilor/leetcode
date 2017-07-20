package sell

func Sell(prices []int) int {
	maxCur := 0
	maxSofar := 0
	for i := 1; i < len(prices); i++ {
		maxCur += prices[i] - prices[i-1]
		maxCur = max(0, maxCur)
		maxSofar = max(maxCur, maxSofar)
	}
	return maxSofar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
