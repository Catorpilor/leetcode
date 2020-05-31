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

// useMath time complexity O(MlgM + NlgN), space complexity O(1)
func useMath(w, h int, horizontalCuts, verticalCuts []int) int {
	// we just need to find the max diffs between the cuts.
	// for example w = 4, h = 4, hc = [1,2,4], vc = [1,3]
	// diff = cur - hc[j-1]
	// the diffs between hc and cake boundary is [1, 1, 2, 1],  max height = 2
	// same approach vc is [1,2,1] max width is 2
	// max area is 4
	m, n := len(horizontalCuts), len(verticalCuts)
	sort.Slice(horizontalCuts, func(i, j int) bool {
		return horizontalCuts[i] < horizontalCuts[j]
	})
	sort.Slice(verticalCuts, func(i, j int) bool {
		return verticalCuts[i] < verticalCuts[j]
	})
	maxW := horizontalCuts[0]
	maxH := verticalCuts[0]
	for i := 1; i <= m; i++ {
		cur := h
		if i < m {
			cur = horizontalCuts[i]
		}
		maxW = utils.Max(maxW, cur-horizontalCuts[i-1])
	}
	for j := 1; j <= n; j++ {
		cur := w
		if j < n {
			cur = verticalCuts[j]
		}
		maxH = utils.Max(maxH, cur-verticalCuts[j-1])
	}
	area := int64(maxW * maxH)
	area %= mod
	return int(area)
}
