package islands

func NumberOfIslands(m, n int, pos [][]int) []int {
	if m*n == 0 || len(pos) == 0 {
		return nil
	}
	ids := make([]int, m*n)
	for i := range ids {
		ids[i] = -1
	}
	wz := make([]int, m*n)
	ret := make([]int, 0, len(pos))
	// hm := make(map[int]bool)
	root := func(p int) int {
		for p != ids[p] {
			ids[p] = ids[ids[p]]
			p = ids[p]
		}
		return p
	}
	union := func(p, q int) bool {
		i, j := root(p), root(q)
		if i == j {
			return false
		}
		if wz[i] < wz[j] {
			ids[i] = j
			wz[j] += wz[i]
		} else {
			ids[j] = i
			wz[i] += wz[j]
		}
		return true
	}
	dirs := [4][2]int{[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, -1}, [2]int{0, 1}}
	var count int
	for i := range pos {
		r, c := pos[i][0], pos[i][1]
		t := r*n + c
		ids[t] = t
		wz[t] = 1
		count++
		for _, v := range dirs {
			x, y := r+v[0], c+v[1]
			tp := x*n + y
			if x >= 0 && x < m && y >= 0 && y < n && ids[tp] != -1 {
				if union(t, tp) {
					count--
				}
			}
		}
		ret = append(ret, count)
	}
	return ret
}
