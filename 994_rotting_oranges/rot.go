package rot

func orangeRotting(grid [][]int) int {
	return useBFS(grid)
}

// useBFS time complexity O(MN), space complexity O(MN)
func useBFS(grid [][]int) int {
	m := len(grid)
	if m < 1 {
		return 0
	}
	n := len(grid[0])
	if n < 1 {
		return 0
	}
	st := make([][]int, 0, m*n)
	freshOnes := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				st = append(st, []int{i, j})
			}
			if grid[i][j] == 1 {
				freshOnes++
			}
		}
	}
	if freshOnes == 0 {
		return 0
	}
	if len(st) == 0 {
		return -1
	}
	var ans int
	dir := [5]int{-1, 0, 1, 0, -1}
	for len(st) > 0 {
		// first level
		curLen := len(st)
		hasFresh := false
		for i := 0; i < curLen; i++ {
			x, y := st[i][0], st[i][1]
			for k := 0; k < 4; k++ {
				nx, ny := x+dir[k], y+dir[k+1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
					st = append(st, []int{nx, ny})
					grid[nx][ny] = 2
					freshOnes--
					hasFresh = true
				}
			}
		}
		if hasFresh {
			ans++
		}
		st = st[curLen:]
	}
	if freshOnes != 0 {
		return -1
	}
	return ans
}
