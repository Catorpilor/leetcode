package lipm

import (
	"github.com/catorpilor/leetcode/utils"
)

func longestIncPath(matrix [][]int) int {
	return dfs(matrix)
}

func dfs(matrix [][]int) int {
	var res int
	m := len(matrix)
	if m < 1 {
		return res
	}
	n := len(matrix)
	if n < 1 {
		return res
	}
	dirs := [5]int{-1, 0, 1, 0, -1}
	// visited := make([]bool, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			out := helper(matrix, dirs, i, j, m, n)
			if out > res {
				res = out
			}
		}
	}
	return res
}

func helper(matrix [][]int, dirs [5]int, x, y, m, n int) int {
	res := 1
	for i := 0; i < 4; i++ {
		nx, ny := x+dirs[i], y+dirs[i+1]
		if nx < 0 || nx >= m || ny < 0 || ny >= n || matrix[nx][ny] >= matrix[x][y] {
			continue
		}
		res = utils.Max(res, 1+helper(matrix, dirs, nx, ny, m, n))
	}
	return res
}
