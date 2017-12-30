package islands

import (
	"github.com/catorpilor/leetcode/utils"
)

func NumOfIslands(grid [][]byte) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	if col == 0 {
		return 0
	}
	var ret int
	// O(M*N) time space
	// Space: Worst case O(M*N)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				// root of tree
				ret++
				dfs(&grid, i, j, row, col)
			}
		}
	}
	return ret
}

func dfs(grid *[][]byte, curRow, curCol, rows, cols int) {
	// set grid[curRow][curCol] to '0' marked visited
	(*grid)[curRow][curCol] = '0'
	if curRow-1 >= 0 && (*grid)[curRow-1][curCol] == '1' {
		dfs(grid, curRow-1, curCol, rows, cols)
	}
	if curRow+1 < rows && (*grid)[curRow+1][curCol] == '1' {
		dfs(grid, curRow+1, curCol, rows, cols)
	}
	if curCol+1 < cols && (*grid)[curRow][curCol+1] == '1' {
		dfs(grid, curRow, curCol+1, rows, cols)
	}
	if curCol-1 >= 0 && (*grid)[curRow][curCol-1] == '1' {
		dfs(grid, curRow, curCol-1, rows, cols)
	}
}

func NumOfIslands2(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}
	// bfs
	var ret int
	type pair struct{ r, c int }
	q := utils.NewQueue()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				ret++
				grid[i][j] = '0' // mark visited
				q.Enroll(pair{i, j})
				for !q.IsEmpty() {
					p := q.Pull().(pair)
					if p.r-1 >= 0 && grid[p.r-1][p.c] == '1' {
						q.Enroll(pair{p.r - 1, p.c})
						grid[p.r-1][p.c] = '0'
					}
					if p.r+1 < rows && grid[p.r+1][p.c] == '1' {
						q.Enroll(pair{p.r + 1, p.c})
						grid[p.r+1][p.c] = '0'
					}
					if p.c-1 >= 0 && grid[p.r][p.c-1] == '1' {
						q.Enroll(pair{p.r, p.c - 1})
						grid[p.r][p.c-1] = '0'
					}
					if p.c+1 < cols && grid[p.r][p.c+1] == '1' {
						q.Enroll(pair{p.r, p.c + 1})
						grid[p.r][p.c+1] = '0'
					}
				}
			}
		}
	}
	return ret
}

func NumOfIslands3(grid [][]byte) int {
	// union-find
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}
	var count int
	ids := make([]int, rows*cols)
	sz := make([]int, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				ids[i*cols+j] = i*cols + j
				sz[i*cols+j] = 1
				count++
			}
		}
	}

	root := func(p int) int {
		for p != ids[p] {
			ids[p] = ids[ids[p]]
			p = ids[p]
		}
		return p
	}
	union := func(p, q int) {
		i, j := root(p), root(q)
		if i == j {
			return
		}
		if sz[i] < sz[j] {
			ids[i] = j
			sz[j] += sz[i]
		} else {
			ids[j] = i
			sz[i] += sz[j]
		}
		count--
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0' // mark visited
				if i > 0 && grid[i-1][j] == '1' {
					union(i*cols+j, (i-1)*cols+j)
				}
				if i+1 < rows && grid[i+1][j] == '1' {
					union(i*cols+j, (i+1)*cols+j)
				}
				if j > 0 && grid[i][j-1] == '1' {
					union(i*cols+j, i*cols+j-1)
				}
				if j+1 < cols && grid[i][j+1] == '1' {
					union(i*cols+j, i*cols+j+1)
				}
			}
		}
	}

	return count
}
