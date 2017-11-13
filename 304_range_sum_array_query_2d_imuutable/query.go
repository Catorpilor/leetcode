package query

// type NumMatrix struct {
// 	accRowMatrix [][]int
// 	row, col     int
// }

// func Constructor(matrix [][]int) NumMatrix {
// 	m := len(matrix)
// 	if m == 0 {
// 		return NumMatrix{}
// 	}
// 	n := len(matrix[0])
// 	if n == 0 {
// 		return NumMatrix{}
// 	}
// 	acculated := make([][]int, m)
// 	for i := range acculated {
// 		acculated[i] = make([]int, n)
// 		acculated[i][0] = matrix[i][0]
// 		for j := 1; j < n; j++ {
// 			acculated[i][j] = acculated[i][j-1] + matrix[i][j]
// 		}
// 	}
// 	return NumMatrix{accRowMatrix: acculated,
// 		row: m,
// 		col: n,
// 	}
// }

type NumMatrix struct {
	sum      [][]int
	row, col int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{}
	}
	n := len(matrix[0])
	if n == 0 {
		return NumMatrix{}
	}
	sum := make([][]int, m+1)
	for i := range sum {
		sum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	// for i := range acculated {
	// 	acculated[i] = make([]int, n+1)
	// 	acculated[i][0] = matrix[i][0]
	// 	for j := 1; j < n; j++ {
	// 		acculated[i][j] = acculated[i][j-1] + matrix[i][j]
	// 	}
	// }
	return NumMatrix{sum: sum,
		row: m,
		col: n,
	}
}

func (nm *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	if row1 > nm.row || row2 < 0 || col2 < 0 || col1 > nm.col {
		return 0
	}
	// var ret int
	// for i := row1; i <= row2; i++ {
	// 	if col1 == 0 {
	// 		ret += nm.accRowMatrix[i][col2]
	// 	} else {
	// 		ret += nm.accRowMatrix[i][col2] - nm.accRowMatrix[i][col1-1]
	// 	}
	// }
	// return ret
	return nm.sum[row2+1][col2+1] - nm.sum[row2+1][col1] - nm.sum[row1][col2+1] + nm.sum[row1][col1]
}
