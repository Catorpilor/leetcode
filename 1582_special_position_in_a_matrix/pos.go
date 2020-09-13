package pos

func specialPos(mat [][]int) int {
	return useStraight(mat)
}

// useStraight time complexity O(MN), space complexity O(M+N)
func useStraight(mat [][]int) int {
	m := len(mat)
	if m < 1 {
		return 0
	}
	n := len(mat[0])
	if n < 1 {
		return 0
	}
	var ans int
	rowCount := make([]int, m)
	colCount := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				rowCount[i]++
				colCount[j]++
			}
		}
	}
	rows := make([]int, 0, m)
	for i := range rowCount {
		if rowCount[i] == 1 {
			rows = append(rows, i)
		}
	}
	for _, i := range rows {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && colCount[j] == 1 {
				ans++
			}
		}
	}
	return ans
}
