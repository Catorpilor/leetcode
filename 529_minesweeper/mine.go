package mine

import "fmt"

func UpdateBoard(board [][]byte, click []int) [][]byte {
	m := len(board)
	if m < 1 {
		return board
	}
	n := len(board[0])
	if n < 1 {
		return board
	}
	if len(click) != 2 {
		return board
	}
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		pos := [][]int{[]int{-1, 0}, []int{-1, 1}, []int{-1, -1}, []int{0, -1}, []int{0, 1}, []int{1, 1}, []int{1, -1}}
		dfs(board, x, y, m, n, pos)
	}
	return board

}

func dfs(board [][]byte, x, y, m, n int, dirs [][]int) {
	if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'E' {
		return
	}
	adjMineCnt := adjMines(board, x, y, m, n)
	if adjMineCnt > 0 {
		board[x][y] = '0' + 1
	} else {
		board[x][y] = byte('B')
		fmt.Printf("board[x]= %v\n", board[x])
		for _, v := range dirs {
			dfs(board, x+v[0], y+v[1], m, n, dirs)
		}
	}

}

func adjMines(board [][]byte, x, y, m, n int) int {
	var cnt int
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i >= 0 && i < m && j >= 0 && j < n && board[i][j] == 'M' {
				cnt++
			}
		}
	}
	return cnt
}
