package tower

import "math"

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

// useDP time complexity O(query_row^2), space complexity O(query_row)
func useDP(poured, query_row, query_glass int) float64 {
	res := make([]float64, query_row+1)
	res[0] = float64(poured)
	for i := 0; i < query_row; i++ {
		for j := i; j >= 0; j-- {
			diff := math.Max(0.0, res[j]-1.0)
			if j+1 <= query_row {
				res[j+1] += diff / 2
			}
			res[j] = diff / 2
		}
	}
	if res[query_glass] >= 1.0 {
		return 1.0
	}
	return res[query_glass]
}
