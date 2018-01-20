package maze

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func HasPath(maze [][]int, start, dest []int) int {
	m := len(maze)
	if m == 0 {
		return -1
	}
	n := len(maze[0])
	if n == 0 {
		return -1
	}
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	dirs := [4][2]int{[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0}}
	ans[start[0]][start[1]] = 1
	dfs(maze, start, dest, &ans, dirs)
	return ans[dest[0]][dest[1]] - 1
	// return dfs2(maze, start, dest, &visited, dirs)

}

func dfs(maze [][]int, start, dest []int, ans *[][]int, dirs [4][2]int) {
	if start[0] == dest[0] && start[1] == dest[1] {
		return
	}
	for _, c := range dirs {
		x, y := start[0], start[1]
		l := (*ans)[x][y]
		for x+c[0] >= 0 && x+c[0] < len(maze) && y+c[1] >= 0 && y+c[1] < len(maze[0]) && maze[x+c[0]][y+c[1]] == 0 {
			x, y, l = x+c[0], y+c[1], l+1
		}
		if (*ans)[x][y] > 0 && l >= (*ans)[x][y] {
			continue
		}
		(*ans)[x][y] = l
		dfs(maze, []int{x, y}, dest, ans, dirs)
	}
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

func HasPath2(maze [][]int, start, dest []int) int {
	m := len(maze)
	if m == 0 {
		return -1
	}
	n := len(maze[0])
	if n == 0 {
		return -1
	}
	q := utils.NewQueue()
	// dp[i][j] represents the minimal length
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	type pair struct {
		pos    []int
		length int
	}
	q.Enroll(pair{pos: start, length: 0})
	dirs := [4][2]int{[2]int{1, 0}, [2]int{-1, 0}, [2]int{0, 1}, [2]int{0, -1}}
	var x, y, l int
	for !q.IsEmpty() {
		cur := q.Pull().(pair)
		for _, c := range dirs {
			x, y, l = cur.pos[0], cur.pos[1], cur.length
			for x >= 0 && x < m && y >= 0 && y < n && maze[x][y] == 0 {
				x, y, l = x+c[0], y+c[1], l+1
			}
			x, y, l = x-c[0], y-c[1], l-1
			if l > dp[x][y] {
				// no need to caculate
				continue
			}
			if l < dp[x][y] {
				dp[x][y] = l
				q.Enroll(pair{pos: []int{x, y}, length: l})
			}
			// put in queue
			// q.Enroll([]int{x, y})
		}
	}
	if dp[dest[0]][dest[1]] == math.MaxInt32 {
		return -1
	}
	return dp[dest[0]][dest[1]]
}
