package noi

import "sort"

func minErase(intervals [][]int) int {
	return useGreedy(intervals)
}

// useGreedy time complexity O(N*lgN) space complexity O(1)
func useGreedy(intervals [][]int) int {
	n := len(intervals)
	if n <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] <= intervals[j][1]
	})
	count := 1
	end := intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
			count++
		}
	}
	return n - count
}
