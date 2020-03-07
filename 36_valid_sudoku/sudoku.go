package sudoku

import "fmt"

func isValid(board [][]byte) bool {
	return useHashMap(board)
}

// useHashMap time complexity O(1), space complexity O(1)
func useHashMap(board [][]byte) bool {
	// rows, cols, boxes are array of hashmaps to do the validation
	rows, cols, boxes := make([]map[byte]bool, 9), make([]map[byte]bool, 9), make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
		boxes[i] = make(map[byte]bool, 9)
	}
	// we spilit the board into 9 boxes boxid = (i/3)*3 + j/3
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			boxid := (i/3)*3 + j/3
			if c != '.' {
				if rows[i][c] || cols[j][c] || boxes[boxid][c] {
					// already exists
					return false
				}
				rows[i][c] = true
				cols[j][c] = true
				boxes[boxid][c] = true
			}
		}
	}
	return true
}

// useHashMapMoreHumanReadable time complexity O(1), space complexity O(1)
func useHashMapMoreHumanReadable(board [][]byte) bool {
	seen := make(map[string]bool)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c != '.' {
				rowkey := fmt.Sprintf("%d in row %d", c, i)
				colkey := fmt.Sprintf("%d in col %d", c, j)
				boxkey := fmt.Sprintf("%d in box %d-%d", c, i/3, j/3)
				if seen[rowkey] || seen[colkey] || seen[boxkey] {
					return false
				}
				seen[rowkey] = true
				seen[colkey] = true
				seen[boxkey] = true
			}
		}
	}
	return true
}
