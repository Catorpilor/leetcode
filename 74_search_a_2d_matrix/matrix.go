package matrix

func searchMatrix(matrix [][]int, target int) bool {
	return useBS(matrix, target)
}

// useBS time complexity O(log(m*n)), space complexity O(1)
func useBS(matrix [][]int, target int) bool {
	m := len(matrix)
	if m < 1 {
		return false
	}
	n := len(matrix[0])
	if n < 1 {
		return false
	}
	l, r := 0, m*n-1
	for l <= r {
		mid := l + (r-l)>>1
		nr, nc := mid/n, mid%n
		if matrix[nr][nc] == target {
			return true
		} else if matrix[nr][nc] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
