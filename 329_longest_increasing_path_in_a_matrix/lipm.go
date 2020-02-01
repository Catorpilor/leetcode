package lipm

import (
	"github.com/catorpilor/leetcode/utils"
)

func longestIncPath(matrix [][]int) int {
	return dfsWithMemo(matrix)
}

// dfs got TLE

// dfsWithMemo time complexity is O(MN), space complexity O(MN)
func dfsWithMemo(matrix [][]int) int {
	var res int
	m := len(matrix)
	if m < 1 {
		return res
	}
	n := len(matrix[0])
	if n < 1 {
		return res
	}
	dirs := [5]int{-1, 0, 1, 0, -1}
	memo := make([]int, m*n)
	// visited := make([]bool, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			out := helper(matrix, memo, dirs, i, j, m, n)
			if out > res {
				res = out
			}
		}
	}
	return res
}

func helper(matrix [][]int, memo []int, dirs [5]int, x, y, m, n int) int {
	res := 1
	// fmt.Printf("cur:(%d, %d)\n", x, y)
	pos := x*n + y
	if memo[pos] != 0 {
		return memo[pos]
	}
	for i := 0; i < 4; i++ {
		nx, ny := x+dirs[i], y+dirs[i+1]
		if nx < 0 || nx >= m || ny < 0 || ny >= n || matrix[nx][ny] >= matrix[x][y] {
			continue
		}
		res = utils.Max(res, 1+helper(matrix, memo, dirs, nx, ny, m, n))
		// fmt.Printf("prev:(%d,%d), cur:(%d, %d), res:%d\n", x, y, nx, ny, res)
	}
	memo[pos] = res
	return res
}
