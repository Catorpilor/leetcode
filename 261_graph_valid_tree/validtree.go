package validtree

func ValidTree(n int, edges [][]int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n > 0 && len(edges) == 0 {
		return false
	}
	// weighted union-find with path compression
	ids, wz := make([]int, n), make([]int, n)
	for i := range ids {
		ids[i], wz[i] = i, 1
	}
	root := func(p int) int {
		for p != ids[p] {
			ids[p] = ids[ids[p]]
			p = ids[p]
		}
		return p
	}
	union := func(i, j int) {
		if wz[i] < wz[j] {
			ids[i] = j
			wz[j] += wz[i]
		} else {
			ids[j] = i
			wz[i] += wz[j]
		}
	}
	for _, v := range edges {
		x, y := root(v[0]), root(v[1])
		if x == y {
			return false
		}
		union(x, y)
		n--
	}

	return n == 1
}
