package inter

import "github.com/catorpilor/leetcode/utils"

func insertIntervals(intervals [][]int, newInterval []int) [][]int {
	return useStraight(intervals, newInterval)
}

// uesStraight time complexity O(N), space complexity O(1)
func useStraight(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	var idx int
	ans := make([][]int, 0, n+1)
	for idx < n && intervals[idx][1] < newInterval[0] {
		ans = append(ans, intervals[idx])
		idx++
	}
	for idx < n && intervals[idx][0] <= newInterval[1] {
		newInterval = []int{utils.Min(newInterval[0], intervals[idx][0]), utils.Max(newInterval[1], intervals[idx][1])}
		idx++
	}
	ans = append(ans, newInterval)
	if idx < n {
		ans = append(ans, intervals[idx:]...)
	}
	return ans
}
