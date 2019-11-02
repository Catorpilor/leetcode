package sudoku

import "fmt"

const (
	n = 3
	N = 3 * 3
)

var (
	resolved bool
)

func solveSudoku(board [][]byte) [][]byte {
	// box size 3
	row := make([][]int, N)
	columns := make([][]int, N)
	box := make([][]int, N)
	res := make([][]byte, N)
	for i := 0; i < N; i++ {
		row[i] = make([]int, N+1)
		columns[i] = make([]int, N+1)
		box[i] = make([]int, N+1)
		res[i] = make([]byte, N)
		copy(res[i], board[i])
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] != '.' {
				placeNumberAtPos(res, i, j, int(board[i][j]-'0'), row, columns, box)
			}
		}
	}
	fmt.Printf("row: %v\n, column: %v\n, box: %v\n", row, columns, box)
	permute(res, 0, 0, row, columns, box)
	return res
}

func placeNumberAtPos(board [][]byte, i, j, num int, row, columns, box [][]int) {
	boxIdx := (i/n)*n + j/n
	(row)[i][num]++
	(columns)[j][num]++
	(box)[boxIdx][num]++
	(board)[i][j] = byte('0' + num)
}

func removeNumberAtPos(board [][]byte, i, j, num int, row, columns, box [][]int) {
	boxIdx := (i/n)*n + j/n
	row[i][num]++
	columns[j][num]++
	box[boxIdx][num]++
	board[i][j] = '.'
}

func isValidPosForNum(i, j, num int, row, columns, box [][]int) bool {
	boxIdx := (i/n)*n + j/n
	return row[i][num]+columns[j][num]+box[boxIdx][num] == 0
}

func permute(board [][]byte, i, j int, row, column, box [][]int) {
	if board[i][j] == '.' {
		for k := 1; k <= N; k++ {
			if isValidPosForNum(i, j, k, row, column, box) {
				placeNumberAtPos(board, i, j, k, row, column, box)
				fmt.Printf("place k:%d at row: %d and col:%d, row[i][k]= %d, col[j][k] = %d and box[boxidx][k] = %d\n", k, i, j, row[i][k], column[j][k],
					box[(i/n)*n+j/n][k])
				placeNext(board, i, j, row, column, box)
				fmt.Printf("place next then  k:%d at row: %d and col:%d, row[i][k]= %d, col[j][k] = %d and box[boxidx][k] = %d\n", k, i, j, row[i][k], column[j][k],
					box[(i/n)*n+j/n][k])
				if !resolved {
					removeNumberAtPos(board, i, j, k, row, column, box)
				}
			}
		}
	} else {
		placeNext(board, i, j, row, column, box)
	}
}

func placeNext(board [][]byte, i, j int, row, column, box [][]int) {
	if i == N-1 && j == N-1 {
		resolved = true
	}
	fmt.Printf("board: %v\n, row: %v \n, column: %v\n, box: %v\n", board, row, column, box)
	if j == N-1 {
		fmt.Println("next row")
		permute(board, i+1, 0, row, column, box)
	} else {
		fmt.Println("next column")
		permute(board, i, j+1, row, column, box)
	}
}

func solveSudoku2(board [][]byte) [][]byte {
	if len(board) == 0 {
		return board
	}
	solve(board)
	return board
}

func solve(board [][]byte) bool {
	var c byte
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				for c = '1'; c <= '9'; c++ {
					if isValid(board, i, j, c) {
						board[i][j] = c
						if solve(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}

		}
	}
	return true
}

func isValid(board [][]byte, row int, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] != '.' && board[i][col] == c {
			return false
		}
		if board[row][i] != '.' && board[row][i] == c {
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i%3] != '.' && board[3*(row/3)+i/3][3*(col/3)+i%3] == c {
			return false
		}
	}
	return true
}
