package boat

import (
	"sort"
)

func minRes(ppl []int, limit int) int {
	return useSort(ppl, limit)
}

// useSort time complexity O(N*log(N)), space complexity O(1)
func useSort(ppl []int, limit int) int {
	sort.Ints(ppl)
	l, r := 0, len(ppl)-1
	var ans int
	for l <= r {
		if ppl[l]+ppl[r] <= limit {
			l++
		}
		r--
		ans++
	}
	return ans
}
