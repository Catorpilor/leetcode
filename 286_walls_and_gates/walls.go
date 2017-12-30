package walls

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func WallsAndGates(rooms [][]int) [][]int {
	rows := len(rooms)
	if rows == 0 {
		return rooms
	}
	cols := len(rooms[0])
	if cols == 0 {
		return rooms
	}
	// dfs version
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if rooms[i][j] == 0 {
				// a gate or a root
				helper(&rooms, i, j, 0, rows, cols) // current deps = 0
			}
		}
	}
	return rooms
}

func helper(rooms *[][]int, i, j, deps, rows, cols int) {
	if i < 0 || j < 0 || i >= rows || j >= cols || (*rooms)[i][j] < deps {
		return
	}
	(*rooms)[i][j] = deps
	helper(rooms, i+1, j, deps+1, rows, cols)
	helper(rooms, i-1, j, deps+1, rows, cols)
	helper(rooms, i, j+1, deps+1, rows, cols)
	helper(rooms, i, j-1, deps+1, rows, cols)
}

const INF = math.MaxInt32

func WallsAndGates2(rooms [][]int) [][]int {
	rows := len(rooms)
	if rows == 0 {
		return rooms
	}
	cols := len(rooms[0])
	if cols == 0 {
		return rooms
	}
	q := utils.NewQueue()
	type pair struct{ i, j int }
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if rooms[i][j] == 0 {
				q.Enroll(pair{i, j})
			}
		}
	}
	// like level order travesal
	for !q.IsEmpty() {
		v := q.Pull().(pair)
		if v.i > 0 && rooms[v.i-1][v.j] == INF {
			rooms[v.i-1][v.j] = rooms[v.i][v.j] + 1
			q.Enroll(pair{v.i - 1, v.j})
		}
		if v.i+1 < rows && rooms[v.i+1][v.j] == INF {
			rooms[v.i+1][v.j] = rooms[v.i][v.j] + 1
			q.Enroll(pair{v.i + 1, v.j})
		}
		if v.j+1 < cols && rooms[v.i][v.j+1] == INF {
			rooms[v.i][v.j+1] = rooms[v.i][v.j] + 1
			q.Enroll(pair{v.i, v.j + 1})
		}
		if v.j > 0 && rooms[v.i][v.j-1] == INF {
			rooms[v.i][v.j-1] = rooms[v.i][v.j] + 1
			q.Enroll(pair{v.i, v.j - 1})
		}
	}
	return rooms
}
