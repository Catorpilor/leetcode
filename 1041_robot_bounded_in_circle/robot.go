package robot

func isCircle(commands string) bool {
	return useMath(commands)
}

// useMath time complexity O(N), space compelxity O(1)
func useMath(commands string) bool {
	d := [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}
	var x, y, i int
	for _, c := range commands {
		if c == 'R' {
			i = (i + 1) % 4
		} else if c == 'L' {
			i = (i + 3) % 4
		} else {
			x, y = x+d[i][0], y+d[i][1]
		}
	}
	return x == 0 && y == 0 || i > 0
}
