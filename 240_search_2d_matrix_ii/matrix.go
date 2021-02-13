package matrix

func searchMatrix(matrix [][]int, target int) bool {
	return useBF(matrix, target)
}

// useBF time complexity O(M+N), space complexity O(1)
func useBF(matrix [][]int, target int) bool {
	m := len(matrix)
	if m < 1 {
		return false
	}
	n := len(matrix[0])
	if n < 1 {
		return false
	}
	row, col := 0, n-1
	for row < m && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
