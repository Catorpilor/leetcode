package box

import (
	"github.com/catorpilor/leetcode/utils"
)

func RemoveBoxes(boxes []int) int {
	n := len(boxes)
	// ans[l][r][k] represents the largest number we can get from boxes[l:r] with k same colored boxes.
	// for example ans[l][r][3] represents the solution for [b_l, ... b_r, b_r, b_r, b_r] (format ABDAA) => (BDAAA)
	// The transition function is to find the maximum among all b_i==b_r for i=l,...,r-1:
	ans := [101][101][101]int{}
	// k >= 1
	return dfs(boxes, &ans, 0, n-1, 1)
}

func dfs(boxes []int, ans *[101][101][101]int, l, r, k int) int {
	if l > r {
		return 0
	}
	if (*ans)[l][r][k] != 0 {
		return (*ans)[l][r][k]
	}
	for r > l && boxes[r] == boxes[r-1] {
		// move r as left as it can
		r, k = r-1, k+1
	}
	for l < r && boxes[l] == boxes[r] {
		l, k = l+1, k+1
	}
	(*ans)[l][r][k] = dfs(boxes, ans, l, r-1, 1) + k*k
	for i := l; i < r; i++ {
		if boxes[i] == boxes[r] {
			(*ans)[l][r][k] = utils.Max((*ans)[l][r][k], dfs(boxes, ans, l, i, k+1)+dfs(boxes, ans, i+1, r-1, 1))
		}
	}
	return (*ans)[l][r][k]
}
