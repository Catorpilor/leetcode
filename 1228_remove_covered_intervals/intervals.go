package intervals

import "sort"

func coveredIntervals(intervals [][]int) int {
	return useSorting(intervals)
}

// useSorting time complexity O(N*logN), space complexity O(1)
func useSorting(intervals [][]int) int {
	// sort intervals based on the start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	n := len(intervals)
	if n <= 1 {
		return n
	}
	left, right := intervals[0][0], intervals[0][1]
	var ans int
	for i := 1; i < n; i++ {
		if intervals[i][0] >= left && intervals[i][1] <= right {
			ans++
		} else {
			if intervals[i][0] <= left && intervals[i][1] >= right {
				ans++
			}
			left, right = intervals[i][0], intervals[i][1]
		}
	}
	return n - ans

}
