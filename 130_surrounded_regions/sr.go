package sr

import "github.com/catorpilor/leetcode/utils/uf"

func solved(board [][]byte) [][]byte {
	return dfs(board)
}

func dfs(board [][]byte) [][]byte {
	m := len(board)
	if m < 1 {
		return board
	}
	n := len(board[0])
	if n < 1 {
		return board
	}
	// copy one
	tmp := make([][]byte, m)
	for i := range tmp {
		tmp[i] = make([]byte, n)
		copy(tmp[i], board[i])
	}
	type pair struct {
		x, y int
	}
	// edges stores the o's position at the boarder
	edges := make([]pair, 0, 2*m+2*n)
	for j := 0; j < n; j++ {
		if tmp[0][j] == 'O' {
			edges = append(edges, pair{0, j})
		}
		if tmp[m-1][j] == 'O' {
			edges = append(edges, pair{m - 1, j})
		}
	}
	for i := 1; i < m-1; i++ {
		if tmp[i][0] == 'O' {
			edges = append(edges, pair{i, 0})
		}
		if tmp[i][n-1] == 'O' {
			edges = append(edges, pair{i, n - 1})
		}
	}
	// mark phase, mark all edge connected 'O's to 'E'
	for _, p := range edges {
		if tmp[p.x][p.y] == 'O' {
			helper(tmp, p.x, p.y, m, n)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if tmp[i][j] == 'O' {
				tmp[i][j] = 'X'
			}
			if tmp[i][j] == 'E' {
				tmp[i][j] = 'O'
			}
		}
	}

	return tmp
}

func helper(board [][]byte, i, j, m, n int) {
	if board[i][j] != 'O' {
		return
	}
	board[i][j] = 'E'
	if i > 0 {
		helper(board, i-1, j, m, n)
	}
	if i < m-1 {
		helper(board, i+1, j, m, n)
	}
	if j > 0 {
		helper(board, i, j-1, m, n)
	}
	if j < n-1 {
		helper(board, i, j+1, m, n)
	}
}

func bfs(board [][]byte) [][]byte {
	m := len(board)
	if m < 1 {
		return board
	}
	n := len(board[0])
	if n < 1 {
		return board
	}
	tmp := make([][]byte, m)
	for i := range tmp {
		tmp[i] = make([]byte, n)
		copy(tmp[i], board[i])
	}
	type pair struct {
		x, y int
	}
	// edges stores the o's position at the boarder
	edges := make([]pair, 0, 2*m+2*n)
	for j := 0; j < n; j++ {
		if tmp[0][j] == 'O' {
			edges = append(edges, pair{0, j})
		}
		if tmp[m-1][j] == 'O' {
			edges = append(edges, pair{m - 1, j})
		}
	}
	for i := 1; i < m-1; i++ {
		if tmp[i][0] == 'O' {
			edges = append(edges, pair{i, 0})
		}
		if tmp[i][n-1] == 'O' {
			edges = append(edges, pair{i, n - 1})
		}
	}
	for len(edges) != 0 {
		next := make([]pair, 0, n+m)
		for _, p := range edges {
			if tmp[p.x][p.y] == 'O' {
				if p.x > 0 && tmp[p.x-1][p.y] == 'O' {
					next = append(next, pair{p.x - 1, p.y})
				}
				if p.y > 0 && tmp[p.x][p.y-1] == 'O' {
					next = append(next, pair{p.x, p.y - 1})
				}
				if p.x < n-1 && tmp[p.x+1][p.y] == 'O' {
					next = append(next, pair{p.x + 1, p.y})
				}
				if p.y < m-1 && tmp[p.x][p.y+1] == 'O' {
					next = append(next, pair{p.x, p.y + 1})
				}
				tmp[p.x][p.y] = 'E'
			}
		}
		edges = next
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if tmp[i][j] == 'O' {
				tmp[i][j] = 'X'
			}
			if tmp[i][j] == 'E' {
				tmp[i][j] = 'O'
			}
		}
	}
	return tmp
}

func withUnionFind(board [][]byte) [][]byte {
	m := len(board)
	if m < 1 {
		return board
	}
	n := len(board[0])
	if n < 1 {
		return board
	}
	tmp := make([][]byte, m)
	for i := range tmp {
		tmp[i] = make([]byte, n)
		copy(tmp[i], board[i])
	}
	dirs := [5]int{-1, 0, 1, 0, -1}
	qf := uf.NewWQUPC(m*n + 1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || j == 0 || i == m-1 || j == n-1) && tmp[i][j] == 'O' {
				// connect to a dummy node
				qf.Union(i*n+j, n*m)
			} else if tmp[i][j] == 'O' {
				for k := 0; k < 4; k++ {
					x, y := dirs[k]+i, dirs[k+1]+j
					if x > 0 && x < m && y > 0 && y < n && tmp[x][y] == 'O' {
						qf.Union(x*n+y, i*n+j)
					}
				}

			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if tmp[i][j] == 'O' && !qf.Find(i*n+j, m*n) {
				tmp[i][j] = 'X'
			}
		}
	}
	return tmp
}
