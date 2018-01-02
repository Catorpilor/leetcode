package island

import "fmt"

func NumOfDistinctIslands(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}

	hm := make(map[string]bool)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				// as a root
				var r []byte
				dfs(&grid, i, j, rows, cols, &r, 'o')
				hm[string(r)] = true
			}
		}
	}
	fmt.Println(hm)
	return len(hm)

}

func dfs(grid *[][]int, i, j, rows, cols int, d *[]byte, ori byte) {
	if i < 0 || j < 0 || i >= rows || j >= cols || (*grid)[i][j] == 0 {
		return
	}
	*d = append((*d), ori)
	(*grid)[i][j] = 0
	dfs(grid, i-1, j, rows, cols, d, 'd')
	dfs(grid, i+1, j, rows, cols, d, 'u')
	dfs(grid, i, j-1, rows, cols, d, 'l')
	dfs(grid, i, j+1, rows, cols, d, 'r')
	*d = append((*d), 'b')
}
