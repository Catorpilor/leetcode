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
	// 0 means no boxes attatched to the left of array at the beginning.
	return dfs(boxes, &ans, 0, n-1, 0)
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
	(*ans)[l][r][k] = dfs(boxes, ans, l, r-1, 0) + (k+1)*(k+1)
	for i := l; i < r; i++ {
		if boxes[i] == boxes[r] {
			(*ans)[l][r][k] = utils.Max((*ans)[l][r][k], dfs(boxes, ans, l, i, k+1)+dfs(boxes, ans, i+1, r-1, 0))
		}
	}
	return (*ans)[l][r][k]
}

func RemoveBoxes2(boxes []int) int {
	// bottom up approach
	n := len(boxes)
	if n <= 1 {
		return n
	}
	ans := [100][100][100]int{}
	for i := 0; i < n; i++ {
		for k := 0; k <= i; k++ {
			ans[i][i][k] = (k + 1) * (k + 1)
		}
	}
	for l := 1; l < n; l++ {
		// l stands for length
		for j := l; j < n; j++ {
			// j stands for the right border
			i := j - l
			for k := 0; k <= i; k++ {
				res := (k+1)*(k+1) + ans[i+1][j][0]
				for m := i + 1; m <= j; m++ {
					if boxes[m] == boxes[j] {
						res = utils.Max(res, ans[i+1][m-1][0]+ans[m][j][k+1])
					}
				}
				ans[i][j][k] = res
			}
		}
	}
	return ans[0][n-1][0]
}
