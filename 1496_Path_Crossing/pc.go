package pc

func pathCrossing(path string) bool {
	return isPathCrossing(path)
}

func isPathCrossing(path string) bool {
	type pos struct {
		x, y int
	}
	set := make(map[pos]bool, len(path))
	start := pos{0, 0}
	set[start] = true
	for i := range path {
		switch path[i] {
		case 'N':
			start.y++
		case 'S':
			start.y--
		case 'E':
			start.x++
		case 'W':
			start.x--
		}
		if set[start] {
			return true
		}
		set[start] = true
	}
	return false
}
