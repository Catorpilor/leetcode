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

func Circles2(grid [][]int) int {
	n := len(grid)
	ret := n
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i // each person i is his/her own circle
	}
	union := func(p, q int) {
		// union person p and q
		// change all id in ids equal to ids[p] to ids[q]
		pid := ids[p]
		for i := range ids {
			if ids[i] == pid {
				ids[i] = ids[q]
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if grid[i][j] == 1 {
				if ids[i] != ids[j] {
					ret--
					union(i, j)
				}
			}
		}
	}
	return ret
}
