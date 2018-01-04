package matrix

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func UpdateMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	if rows == 0 {
		return matrix
	}
	cols := len(matrix[0])
	if cols == 0 {
		return matrix
	}
	ret := make([][]int, rows)
	for i := range ret {
		ret[i] = make([]int, cols)
		copy(ret[i], matrix[i])
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if ret[i][j] == 1 && !hasZeroNeighbor(ret, i, j, rows, cols) {
				// no need to be root
				ret[i][j] = rows + cols + 1
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if ret[i][j] == 1 {
				dfs(&ret, i, j, rows, cols, -1)
			}
		}
	}
	return ret
}

func dfs(res *[][]int, i, j, rows, cols, deps int) {
	if i < 0 || j < 0 || i >= rows || j >= cols || (*res)[i][j] <= deps {
		return
	}
	if deps > 0 {
		(*res)[i][j] = deps
	}
	dfs(res, i-1, j, rows, cols, (*res)[i][j]+1)
	dfs(res, i+1, j, rows, cols, (*res)[i][j]+1)
	dfs(res, i, j-1, rows, cols, (*res)[i][j]+1)
	dfs(res, i, j+1, rows, cols, (*res)[i][j]+1)

}

func hasZeroNeighbor(matrix [][]int, i, j, rows, cols int) bool {
	if i > 0 && matrix[i-1][j] == 0 {
		return true
	}
	if j > 0 && matrix[i][j-1] == 0 {
		return true
	}
	if i < rows-1 && matrix[i+1][j] == 0 {
		return true
	}
	if j < cols-1 && matrix[i][j+1] == 0 {
		return true
	}
	return false

}

func UpdateMatrix2(matrix [][]int) [][]int {
	rows := len(matrix)
	if rows == 0 {
		return matrix
	}
	cols := len(matrix[0])
	if cols == 0 {
		return matrix
	}
	ret := make([][]int, rows)
	for i := range ret {
		ret[i] = make([]int, cols)
		copy(ret[i], matrix[i])
	}
	q := utils.NewQueue()
	type pair struct{ x, y int }
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if ret[i][j] == 0 {
				q.Enroll(pair{i, j})
			} else {
				ret[i][j] = math.MaxInt32
			}
		}
	}
	dirs := [4][2]int{[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, 1}, [2]int{0, -1}}
	for !q.IsEmpty() {
		v := q.Pull().(pair)
		for _, d := range dirs {
			l, r := v.x+d[0], v.y+d[1]
			if l < 0 || r < 0 || l >= rows || r >= cols || ret[l][r] <= ret[v.x][v.y]+1 {
				continue
			}
			q.Enroll(pair{l, r})
			ret[l][r] = ret[v.x][v.y] + 1
		}
	}
	return ret
}

func UpdateMatrix3(matrix [][]int) [][]int {
	rows := len(matrix)
	if rows == 0 {
		return matrix
	}
	cols := len(matrix[0])
	if cols == 0 {
		return matrix
	}
	ret := make([][]int, rows)
	for i := range ret {
		ret[i] = make([]int, cols)
		// copy(ret[i], matrix[i])
	}
	tp := rows * cols
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 1 {
				// left and upcell
				leftC, upperC := tp, tp
				if i > 0 {
					upperC = ret[i-1][j]
				}
				if j > 0 {
					leftC = ret[i][j-1]
				}
				ret[i][j] = utils.Min(upperC, leftC) + 1
			}
		}
	}
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if matrix[i][j] == 1 {
				// right and bottom cell
				rightC, bottomC := tp, tp
				if i < rows-1 {
					bottomC = ret[i+1][j]
				}
				if j < cols-1 {
					rightC = ret[i][j+1]
				}
				ret[i][j] = utils.Min(utils.Min(rightC, bottomC)+1, ret[i][j])
			}
		}
	}
	return ret
}
