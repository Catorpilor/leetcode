package tower

func queryGlass(poured, query_row, query_glass int) float64 {
	return useSimpleMath(poured, query_row, query_glass)
}

// useSimpleMath time complexity O(1), space complexity O(1)
func useSimpleMath(poured, query_row, query_glass int) float64 {
	// at most 100th row
	res := make([][]float64, 101)
	for i := range res {
		res[i] = make([]float64, 101)
	}
	res[0][0] = float64(poured)
	for i := 0; i < 100; i++ {
		for j := 0; j <= i; j++ {
			if res[i][j] >= 1.0 {
				res[i+1][j] += (res[i][j] - 1.0) / 2
				res[i+1][j+1] += (res[i][j] - 1.0) / 2
				res[i][j] = 1.0
			}
		}
	}
	return res[query_row][query_glass]
}
