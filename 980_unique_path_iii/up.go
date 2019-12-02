package up

import "fmt"

func uniquePath(grid [][]int) int {
	var res int
	var sx, sy, ex, ey, cells int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				sx, sy = i, j
			} else if grid[i][j] == 0 {
				cells++
			} else if grid[i][j] == 2 {
				ex, ey = i, j
			}
		}
	}
	// fmt.Printf("ex:%d, ey:%d, sx:%d, sy:%d, cells:%d\n", ex, ey, sx, sy, cells)
	dfs(&res, grid, sx, sy, ex, ey, &cells)
	return res
}

func dfs(res *int, grid [][]int, sx, sy, ex, ey int, empty *int) {
	fmt.Printf("ex:%d, ey:%d, sx:%d, sy:%d, cells:%d\n", ex, ey, sx, sy, *empty)
	if sx < 0 || sx >= len(grid) || sy < 0 || sy >= len(grid[0]) || grid[sx][sy] < 0 {
		return
	}
	if sx == ex && sy == ey {
		if (*empty) == 0 {
			(*res)++
		}
		return
	}
	grid[sx][sy] = -2
	(*empty)--
	dfs(res, grid, sx, sy-1, ex, ey, empty)
	dfs(res, grid, sx, sy+1, ex, ey, empty)
	dfs(res, grid, sx-1, sy, ex, ey, empty)
	dfs(res, grid, sx+1, sy, ex, ey, empty)

	(*empty)++
	grid[sx][sy] = 0
}
