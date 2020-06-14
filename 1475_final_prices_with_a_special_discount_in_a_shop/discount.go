package discount

func finalPrices(prices []int) []int {
	return useStack(prices)
}

// useStack time complexity O(N), space complexity O(N)
func useStack(prices []int) []int {
	n := len(prices)
	if n <= 1 {
		return prices
	}
	st := make([]int, 0, n)
	for i := range prices {
		nst := len(st)
		for nst > 0 && prices[st[nst-1]] >= prices[i] {
			prices[st[nst-1]] -= prices[i]
			st = st[:nst-1]
			nst--
		}
		st = append(st, i)
	}
	return prices
}
