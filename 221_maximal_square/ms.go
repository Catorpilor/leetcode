package ms

import (
	"github.com/catorpilor/leetcode/utils"
)

// MaxSquare returns a max all 1s square inside matrix
func MaxSquare(matrix [][]int) int {
	if matrix == nil {
		return 0
	}
	// brute force
	// if we find a 1, we move diagonally
	// and check this square if it is all 1s
	// if it is true update the maxSqrLen,
	// otherwise move forward
	// time complexity: O((mn)^2)
	// space: O(1)
	row, col := len(matrix), len(matrix[0])
	var maxSqlLen int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 1 {
				tm, flag := 1, true
				for tm+i < row && tm+j < col && flag {
					for k := j; k <= tm+j; k++ {
						if matrix[i+tm][k] == 0 {
							flag = false
							break
						}
					}
					for k := i; k <= i+tm; k++ {
						if matrix[k][j+tm] == 0 {
							flag = false
							break
						}
					}
					if flag {
						tm += 1
					}
				}
				maxSqlLen = utils.Max(maxSqlLen, tm)

			}
		}
	}
	return maxSqlLen * maxSqlLen
}

// MaxSquare2 returns a all 1s square inside matrix
func MaxSquare2(matrix [][]int) int {
	if matrix == nil {
		return 0
	}
	row := len(matrix)
	var col int
	if row > 0 {
		col = len(matrix[0])
	}
	if row == 0 || col == 0 {
		return 0
	}
	var maxSqrLen int
	// allocate a matrix
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		dp[i][0] = matrix[i][0]
		if matrix[i][0] == 1 {
			maxSqrLen = 1
		}
	}
	for j := 0; j < col; j++ {
		dp[0][j] = matrix[0][j]
		if matrix[0][j] == 1 {
			maxSqrLen = 1
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = utils.Min(dp[i-1][j], utils.Min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
			maxSqrLen = utils.Max(maxSqrLen, dp[i][j])
		}
	}
	return maxSqrLen * maxSqrLen
}

// MaxSquare3 returns a max square with 1s inside matrix
func MaxSquare3(matrix [][]int) int {
	if matrix == nil {
		return 0
	}
	row := len(matrix)
	var col int
	if row > 0 {
		col = len(matrix[0])
	}
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([]int, col+1)
	var mx, prev int
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			temp := dp[j]
			if matrix[i-1][j-1] == 1 {
				dp[j] = utils.Min(dp[j], utils.Min(dp[j-1], prev)) + 1
				mx = utils.Max(mx, dp[j])
			}
			prev = temp
		}
	}
	return mx * mx
}
