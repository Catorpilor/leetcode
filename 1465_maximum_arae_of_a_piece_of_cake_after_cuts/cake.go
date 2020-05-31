package cake

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

const (
	mod = 10_000_000_00 + 7
)

func maxArea(w, h int, horizontal, vertical []int) int {
	return useBruteForce(w, h, horizontal, vertical)
}

// useBruteForece time complexity O(MN), space complexity O(1)
// got TLE
func useBruteForce(w, h int, horizontalCuts, verticalCuts []int) int {
	m, n := len(horizontalCuts), len(verticalCuts)
	sort.Slice(horizontalCuts, func(i, j int) bool {
		return horizontalCuts[i] < horizontalCuts[j]
	})
	sort.Slice(verticalCuts, func(i, j int) bool {
		return verticalCuts[i] < verticalCuts[j]
	})
	var ans int
	for i := 0; i <= m; i++ {
		upB, bottomB := 0, h
		if i > 0 {
			upB = horizontalCuts[i-1]
		}
		if i < m {
			bottomB = horizontalCuts[i]
		}
		for j := 0; j <= n; j++ {
			leftB, rightB := 0, w
			if j != 0 {
				leftB = verticalCuts[j-1]
			}
			if j < n {
				rightB = verticalCuts[j]
			}
			area := int64(bottomB-upB) * int64(rightB-leftB)
			// fmt.Printf("upB: %d, rightB:%d, leftB:%d, bottomB:%d, area:%d\n", upB, rightB, leftB, bottomB, area)
			area %= mod
			ans = utils.Max(ans, int(area))
		}
	}
	return ans
}
