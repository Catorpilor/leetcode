package islands

import (
	"github.com/catorpilor/leetcode/utils"
)

// numOfIslands time complexity O(MN), space complexity O(MN)
func numOfIslands(grid [][]byte) int {
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
				dfs(grid, i, j, row, col)
			}
		}
	}
	return ret
}

func dfs(grid [][]byte, i, j, m, n int) {
	// set grid[curRow][curCol] to '0' marked visited
	grid[i][j] = '0'
	pos := [5]int{-1, 0, 1, 0, -1}
	for k := 0; k < 4; k++ {
		nx, ny := i+pos[k], j+pos[k+1]
		if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == '1' {
			dfs(grid, nx, ny, m, n)
		}
	}
}

// useBfs time complexity O(MN), space complexity O(MN)
func useBfs(grid [][]byte) int {
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
	pos := [5]int{-1, 0, 1, 0, -1}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				ret++
				grid[i][j] = '0' // mark visited
				q.Enroll(pair{i, j})
				for !q.IsEmpty() {
					p := q.Pull().(pair)
					for k := 0; k < 4; k++ {
						nx, ny := p.r+pos[k], p.c+pos[k+1]
						if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == '1' {
							q.Enroll(pair{nx, ny})
							grid[nx][ny] = '0'
						}
					}
				}
			}
		}
	}
	return ret
}

// useUnionFind time complexity O(MN), space complexity O(MN)
func useUnionFind(grid [][]byte) int {
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
	pos := [5]int{-1, 0, 1, 0, -1}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0' // mark visited
				for k := 0; k < 4; k++ {
					nx, ny := i+pos[k], j+pos[k+1]
					if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == '1' {
						union(i*cols+j, nx*cols+ny)
					}
				}
				// if i > 0 && grid[i-1][j] == '1' {
				// 	union(i*cols+j, (i-1)*cols+j)
				// }
				// if i+1 < rows && grid[i+1][j] == '1' {
				// 	union(i*cols+j, (i+1)*cols+j)
				// }
				// if j > 0 && grid[i][j-1] == '1' {
				// 	union(i*cols+j, i*cols+j-1)
				// }
				// if j+1 < cols && grid[i][j+1] == '1' {
				// 	union(i*cols+j, i*cols+j+1)
				// }
			}
		}
	}

	return count
}
