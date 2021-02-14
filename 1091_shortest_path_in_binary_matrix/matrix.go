package matrix

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func shortestPath(matrix [][]int) int {
	return -1

}

// useDFS got TLE
func useDFS(grid [][]int) int {
	n := len(grid)
	if n < 1 {
		return -1
	}
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	ans := math.MaxInt32
	visited := make([]bool, n*n)
	helper(grid, *ans, 0, 0, 1, visited, n)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func helper(grid [][]int, ans *int, x, y, k int, visited []bool, n int) {
	if x < 0 || x >= n || y < 0 || y >= n || visited[x*n+y] || grid[x][y] == 1 {
		return
	}
	if x == n-1 && y == n-1 {
		*ans = utils.Min((*ans), k)
		return
	}
	visited[x*n+y] = true
	dirs := [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}, []int{1, 1},
		[]int{-1, 1}, []int{-1, -1}, []int{1, -1}}
	for _, dir := range dirs {
		nx, ny := x+dir[0], y+dir[1]
		helper(grid, ans, nx, ny, k+1, visited, n)
	}
	visited[x*n+y] = false
	return
}
