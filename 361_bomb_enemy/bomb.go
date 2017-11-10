package bomb

import (
	"github.com/catorpilor/leetcode/utils"
)

func MaxKill(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	col := make([]int, n)
	var re, ret int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'W' {
				continue
			}
			if j == 0 || grid[i][j-1] == 'W' {
				// need to caculate enmies in row i
				re = enmyInRow(grid, i, j, n)
			}
			if i == 0 || grid[i-1][j] == 'W' {
				col[j] = enmyInCol(grid, i, j, m)
			}
			if grid[i][j] == '0' {
				ret = utils.Max(ret, col[j]+re)
			}
		}
	}
	return ret
}

func enmyInRow(grid [][]byte, i, j, n int) int {
	var ret int
	for j < n && grid[i][j] != 'W' {
		if grid[i][j] == 'E' {
			ret += 1
		}
		j++
	}
	return ret
}

func enmyInCol(grid [][]byte, i, j, m int) int {
	var ret int
	for i < m && grid[i][j] != 'W' {
		if grid[i][j] == 'E' {
			ret += 1
		}
		i += 1
	}
	return ret
}
