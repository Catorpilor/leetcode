package chips

import "github.com/catorpilor/leetcode/utils"

func minCost(position []int) int {
	return useSimple(position)
}

// useSimple time complexity O(n), space complexity O(1)
func useSimple(position []int) int {
	// because we can move both odds and evens chips with zero cost.
	// in the end there are just two piles.
	cnt := make([]int, 2)
	for _, p := range position {
		if p&1 != 0 {
			cnt[1] += p
		} else {
			cnt[0] += p
		}
	}
	return utils.Min(cnt[0], cnt[1])
}
