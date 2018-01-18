package maze

import (
	"github.com/catorpilor/leetcode/utils"
)

func HasPath(maze [][]int, start, dest []int) bool {
	m := len(maze)
	if m == 0 {
		return false
	}
	n := len(maze[0])
	if n == 0 {
		return false
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	dirs := [4][2]int{[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0}}
	return dfs2(maze, start, dest, &visited, dirs)

}

func dfs(maze [][]int, start, dest []int, visited *[][]bool, dirs [4][2]int) bool {
	x, y := start[0], start[1]
	if (*visited)[x][y] {
		return false
	}
	if x == dest[0] && y == dest[1] {
		return true
	}
	(*visited)[x][y] = true

	// x y represts the point to make desions which direction to go.

	for _, c := range dirs {
		// tempx, tempy = x+c[0], y+c[1]
		newStart := roll(maze, x, y, c[0], c[1])
		if dfs(maze, newStart, dest, visited, dirs) {
			return true
		}
	}
	return false
}

func dfs2(maze [][]int, start, dest []int, visited *[][]bool, dirs [4][2]int) bool {
	if (*visited)[start[0]][start[1]] {
		return false
	}
	if start[0] == dest[0] && start[1] == dest[1] {
		return true
	}
	(*visited)[start[0]][start[1]] = true
	var x, y int
	for _, c := range dirs {
		x, y = start[0], start[1]
		if x+c[0] < len(maze) && x+c[0] >= 0 && y+c[1] < len(maze[0]) && y+c[1] >= 0 && maze[x+c[0]][y+c[1]] == 0 {
			x, y = x+c[0], y+c[1]
		}
		if dfs2(maze, []int{x, y}, dest, visited, dirs) {
			return true
		}
	}
	return false
}

func canRoll(maze [][]int, x, y int) bool {
	if x < 0 || x >= len(maze) || y >= len(maze[0]) || y < 0 {
		// stop at boarder
		return false
	}
	// stop at wall
	return maze[x][y] != 1
}

func roll(maze [][]int, x, y, rowInc, colInc int) []int {
	for canRoll(maze, x+rowInc, y+colInc) {
		x, y = x+rowInc, y+colInc
	}
	return []int{x, y}
}

func HasPath2(maze [][]int, start, dest []int) bool {
	m := len(maze)
	if m == 0 {
		return false
	}
	n := len(maze[0])
	if n == 0 {
		return false
	}
	q := utils.NewQueue()
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	q.Enroll(start)
	visited[start[0]][start[1]] = true
	dirs := [4][2]int{[2]int{1, 0}, [2]int{-1, 0}, [2]int{0, 1}, [2]int{0, -1}}
	var x, y int
	for !q.IsEmpty() {
		cur := q.Pull().([]int)
		for _, c := range dirs {
			x, y = cur[0], cur[1]
			for x >= 0 && x < m && y >= 0 && y < n && maze[x][y] == 0 {
				x, y = x+c[0], y+c[1]
			}
			x, y = x-c[0], y-c[1]
			if visited[x][y] {
				continue
			}
			// mark as visited
			visited[x][y] = true
			// desision point
			if x == dest[0] && y == dest[1] {
				return true
			}
			// put in queue
			q.Enroll([]int{x, y})
		}
	}
	return false
}
