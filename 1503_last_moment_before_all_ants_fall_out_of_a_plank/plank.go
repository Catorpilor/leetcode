package plank

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func lastMoment(n int, left, right []int) int {
	return useMath(n, left, right)
}

// useMath time complexity O(n1Log(n1) + n2log(n2)), space complexity O(1)
func useMath(n int, left, right []int) int {
	nl, nr := len(left), len(right)
	if nl == 0 && nr == 0 {
		return 0
	}
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})
	// we pick the max pos of ants go to left
	// and pick the min pos of ants go to right
	if nl == 0 {
		return n - right[0]
	}
	if nr == 0 {
		return left[nl-1]
	}
	return utils.Max(left[nl-1], n-right[0])
}
