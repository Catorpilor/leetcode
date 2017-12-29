package circle

func Circles(grid [][]int) int {
	rows := len(grid)
	var ret int
	visited := make([]bool, rows)
	for i := 0; i < rows; i++ {
		if !visited[i] {
			dfs(grid, &visited, i, rows)
			ret++
		}
	}
	return ret
}

func dfs(grid [][]int, visited *[]bool, i, n int) {
	for j := 0; j < n; j++ {
		// unvisited person j in the current friend circle
		if !(*visited)[j] && grid[i][j] == 1 {
			(*visited)[j] = true
			dfs(grid, visited, j, n)
		}
	}
}
