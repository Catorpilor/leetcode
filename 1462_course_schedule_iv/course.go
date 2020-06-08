package course

func checkIfPrerequisite(n int, pres [][]int, queries [][]int) []bool {
	return useFloydWarshall(n, pres, queries)
}

// useFloydWarshall time complexity $$O(N^3)$$, space complexity $$O(N^2)$$.
func useFloydWarshall(n int, pres, queries [][]int) []bool {
	connnected := make([][]bool, n)
	for i := range connnected {
		connnected[i] = make([]bool, n)
	}
	for _, p := range pres {
		connnected[p[0]][p[1]] = true
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				connnected[i][j] = connnected[i][j] || (connnected[i][k] && connnected[k][j])
			}
		}
	}
	ans := make([]bool, 0, len(queries))
	for _, q := range queries {
		ans = append(ans, connnected[q[0]][q[1]])
	}
	return ans
}
