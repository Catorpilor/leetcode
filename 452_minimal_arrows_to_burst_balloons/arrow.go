package arrow

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func minArrows(points [][]int) int {
	return useTwoPointers(points)
}

// useTwoPointers time complexity O(N*logN), space complexity O(1)
func useTwoPointers(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return n
	}
	// sort points based the start/end position
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	// boundaries
	left, right := points[0][0], points[0][1]
	ans := 1
	for i := 1; i < n; i++ {
		if points[i][0] > right {
			// update boundaries
			right = points[i][1]
			ans++
		} else {
			right = utils.Min(right, points[i][1])
		}
		_ = left
		left = points[i][0]
	}
	return ans
}

// useBoundary time complexity O(N*logN), space complexity O(1)
func useBoundary(points [][]int) int {
	// since we only care the right boundary
	// so the sorting part can based on the end position.
	n := len(points)
	if n <= 1 {
		return n
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	curB := points[0]
	ans := 1
	for i := 1; i < n; i++ {
		if points[i][0] > curB[1] {
			ans++
			curB = points[i]
		}
	}
	return ans
}
