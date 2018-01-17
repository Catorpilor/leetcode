package maze

func HasPath(maze [][]int, start, dest []int) bool {
	m := len(maze)
	if m == 0 {
		return false
	}
	n := len(maze[0])
	if n == 0 {
		return false
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	dirs := [4][2]int{[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0}}
	return dfs(maze, start, dest, &visited, dirs)

}

func dfs(maze [][]int, start, dest []int, visited *[][]bool, dirs [4][2]int) bool {
	x, y := start[0], start[1]
	if (*visited)[start[0]][start[1]] {
		return false
	}
	if start[0] == dest[0] && start[1] == dest[1] {
		return true
	}
	(*visited)[x][y] = true

	// x y represts the point to make desions which direction to go.

	for _, c := range dirs {
		// tempx, tempy = x+c[0], y+c[1]
		newStart := roll(maze, x, y, c[0], c[1])
		if dfs(maze, newStart, dest, visited, dirs) {
			return true
		}
	}
	return false
}

func canRoll(maze [][]int, x, y int) bool {
	if x < 0 || x >= len(maze) || y >= len(maze[0]) || y < 0 {
		// stop at boarder
		return false
	}
	// stop at wall
	return maze[x][y] != 1
}

func roll(maze [][]int, x, y, rowInc, colInc int) []int {
	for canRoll(maze, x+rowInc, y+colInc) {
		x, y = x+rowInc, y+colInc
	}
	return []int{x, y}
}
