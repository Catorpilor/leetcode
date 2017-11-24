package word

func Exists(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if helper(board, i, j, word, 0, m, n) {
				return true
			}
		}
	}
	return false
}

func helper(board [][]byte, i, j int, word string, l, m, n int) bool {
	if l == len(word) {
		return true
	}
	if i < 0 || j < 0 || i == m || j == n {
		return false
	}
	if board[i][j] != word[l] {
		return false
	}
	board[i][j] ^= 255
	ret := helper(board, i, j+1, word, l+1, m, n) || helper(board, i, j-1, word, l+1, m, n) || helper(board, i+1, j, word, l+1, m, n) || helper(board, i-1, j, word, l+1, m, n)
	board[i][j] ^= 255
	return ret
}
