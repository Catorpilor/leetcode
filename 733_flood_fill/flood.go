package flood

func floodFill(image [][]int, x, y, newColor int) [][]int {
	m := len(image)
	if m < 1 {
		return image
	}
	n := len(image[0])
	if n < 1 {
		return image
	}
	old := image[x][y]
	if old == newColor {
		return image
	}
	useDfs(image, m, n, x, y, newColor, old)
	return image
}

// useDfs time complexity O(N), space complexity O(N)
func useDfs(image [][]int, m, n, x, y, newColor, old int) {
	if x < 0 || x >= m || y < 0 || y >= n || image[x][y] != old {
		return
	}
	image[x][y] = newColor
	dirs := [5]int{1, 0, -1, 0, 1}
	for i := 0; i < 4; i++ {
		nx, ny := x+dirs[i], y+dirs[i+1]
		useDfs(image, m, n, nx, ny, newColor, old)
	}
}
