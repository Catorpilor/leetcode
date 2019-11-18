package ws

import (
	"fmt"

	"github.com/catorpilor/LeetCode/utils"
)

func findWords(board [][]byte, words []string) []string {
	return backTrack(board, words)
}

func backTrack(board [][]byte, words []string) []string {
	var res []string
	n := len(board)
	if n < 1 {
		return res
	}
	m := len(board[0])
	if m < 1 {
		return res
	}
	for _, word := range words {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if helper(&res, word, board, i, j, n, m, 0) {
					goto PP
				}
			}
		}
	PP:
	}
	return res
}

func helper(res *[]string, s string, board [][]byte, x, y, n, m, idx int) bool {
	fmt.Printf("helper s:%s, x:%d, y:%d, idx:%d\n", s, x, y, idx)
	if idx == len(s) {
		*res = append(*res, s)
		fmt.Printf("append to res:%v, s:%s\n", *res, s)
		return true
	}
	if x < 0 || x >= n || y < 0 || y >= m {
		return false
	}
	if board[x][y] != s[idx] {
		return false
	}
	board[x][y] ^= 128
	ret := helper(res, s, board, x, y+1, n, m, idx+1) || helper(res, s, board, x, y-1, n, m, idx+1) ||
		helper(res, s, board, x-1, y, n, m, idx+1) || helper(res, s, board, x+1, y, n, m, idx+1)
	board[x][y] ^= 128
	return ret
}

func withTrie(board [][]byte, words []string) []string {
	var res []string
	n := len(board)
	if n < 1 {
		return res
	}
	m := len(board[0])
	if m < 1 {
		return res
	}
	tst := utils.NewTST()
	for _, w := range words {
		tst.Put(w, w)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			permute(&res, board, tst.Root, i, j, n, m)
		}
	}
	return res
}

func permute(res *[]string, board [][]byte, node *utils.TSTNode, i, j, n, m int) {
	c := board[i][j]
	if c == '#' || node == nil {
		return
	}
	if node.Value != nil {
		if s, ok := node.Value.(string); ok {
			*res = append(*res, s)
			node.Value = nil // remove duplicates
			return
		}
	}
	if c > node.Cb {
		node = node.Right
	} else if c < node.Cb {
		node = node.Left
	} else {
		node = node.Middle
	}
	board[i][j] = '#'
	if i > 0 {
		permute(res, board, node, i-1, j, n, m)
	}
	if j > 0 {
		permute(res, board, node, i, j-1, n, m)
	}
	if i < n-1 {
		permute(res, board, node, i+1, j, n, m)
	}
	if j < m-1 {
		permute(res, board, node, i, j+1, n, m)
	}
	board[i][j] = c
}
