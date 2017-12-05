package matrix

import (
	"github.com/catorpilor/leetcode/utils"
)

func MaxMatrix(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	var ret int
	left, right, height := make([]int, n), make([]int, n), make([]int, n)
	for i := range right {
		right[i] = n
	}
	// height[j] represents the height at matrix[i][j] how many 1s
	// left[j] means the left-boundary of rectangle of height[j]
	// right[j] means the right boundary of rectangle of height[j]
	// for example with input matrix
	//        0 0 0 1 0 0 0
	//        0 0 1 1 1 0 0
	//        0 1 1 1 1 1 0

	// row 0
	// height 0 0 0 1 0 0 0
	// left   0 0 0 3 0 0 0
	// right  7 7 7 4 7 7 7
	// row 1
	// height 0 0 1 2 1 0 0
	// left   0 0 2 3 2 0 0
	// right  7 7 5 4 5 7 7
	// row 2
	// height 0 1 2 3 2 1 0
	// left   0 1 2 3 2 1 0
	// right  7 6 5 4 5 6 7

	cur_left, cur_right := 0, n
	for i := 0; i < m; i++ {
		cur_left, cur_right = 0, n
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				height[j]++
				left[j] = utils.Max(cur_left, left[j])
			} else {
				height[j] = 0
				left[j] = 0
				cur_left = j + 1
			}
		}
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				right[j] = utils.Min(cur_right, right[j])
			} else {
				right[j], cur_right = n, j
			}
		}
		for j := 0; j < n; j++ {
			ret = utils.Max(ret, (right[j]-left[j])*height[j])
		}
	}
	return ret
}
