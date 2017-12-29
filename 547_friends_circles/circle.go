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
	// quick-find,
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

func Circles3(grid [][]int) int {
	n := len(grid)
	ret := n
	// quick-union
	ids := make([]int, n) // ids[i] represents the parent of i
	for i := range ids {
		ids[i] = i
	}
	root := func(i int) int {
		for i != ids[i] {
			i = ids[i]
		}
		return i
	}
	union := func(p, q int) {
		i, j := root(p), root(q)
		ids[i] = j
	}
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 1 {
				if root(i) != root(j) {
					ret--
					union(i, j)
				}
			}
		}
	}
	return ret
}

func Circles4(grid [][]int) int {
	n := len(grid)
	ret := n
	// quick-union with weighted
	ids := make([]int, n) // ids[i] represents the parent of i
	sz := make([]int, n)  // sz[i] reprents the number of elements of root i
	for i := range ids {
		ids[i] = i
		sz[i] = 1
	}
	root := func(i int) int {
		for i != ids[i] {
			i = ids[i]
		}
		return i
	}
	union := func(p, q int) {
		i, j := root(p), root(q)
		// merge smaller tree into larger tree
		if i == j {
			// no need to merge
			return
		}
		if sz[i] < sz[j] {
			ids[i] = j
			sz[j] += sz[i]
		} else {
			ids[j] = i
			sz[i] += sz[j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 1 {
				if root(i) != root(j) {
					ret--
					union(i, j)
				}
			}
		}
	}
	return ret
}

func Circles5(grid [][]int) int {
	n := len(grid)
	ret := n
	// weighted quick-union with path compression
	ids := make([]int, n) // ids[i] represents the parent of i
	sz := make([]int, n)
	for i := range ids {
		ids[i] = i
		sz[i] = 1
	}
	root := func(i int) int {
		for i != ids[i] {
			ids[i] = ids[ids[i]] // make every other node in path point to its grandparent
			i = ids[i]
		}
		return i
	}
	union := func(p, q int) {
		i, j := root(p), root(q)
		if sz[i] < sz[j] {
			ids[i] = j
			sz[j] += sz[i]
		} else {
			ids[j] = i
			sz[i] += sz[j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 1 {
				if root(i) != root(j) {
					ret--
					union(i, j)
				}
			}
		}
	}
	return ret
}
