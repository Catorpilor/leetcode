package dist

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func shortestDistance(s string, c byte) []int {
	return useBinarySearch(s, c)
}

// useBinarySearch time complexity O(N*log(N)), space complexity O(N)
func useBinarySearch(s string, c byte) []int {
	n := len(s)
	pos := make([]int, 0, n)
	for i := range s {
		if s[i] == c {
			pos = append(pos, i)
		}
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = math.MaxInt32
	}
	for i := range s {
		if s[i] != c {
			pp := utils.LowerBound(pos, i)
			if pp == 0 {
				ans[i] = pos[pp] - i
			} else if pp == len(pos) {
				ans[i] = i - pos[pp-1]
			} else {
				ans[i] = utils.Min(i-pos[pp-1], pos[pp]-i)
			}
		} else {
			ans[i] = 0
		}
	}
	return ans
}
