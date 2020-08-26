package coins

import "sort"

func maxCoins(piles []int) int {
	return useSort(piles)
}

// useSort time complexity O(N*log(N)), space complexity O(1)
func useSort(piles []int) int {
	sort.Ints(piles)
	l, r := 0, len(piles)-1
	var ans int
	for l < r {
		ans += piles[r-1]
		l++
		r -= 2
	}
	return ans
}
