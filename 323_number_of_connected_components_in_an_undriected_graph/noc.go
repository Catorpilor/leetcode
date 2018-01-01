package noc

func CountComponents(n int, edges [][]int) int {
	if n <= 1 {
		return n
	}
	// weighted union-find with path compression
	ids := make([]int, n)
	wz := make([]int, n)
	for i := range ids {
		ids[i] = i
		wz[i] = 1
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
	ret := n
	for _, v := range edges {
		r1, r2 := root(v[0]), root(v[1])
		if r1 != r2 {
			union(r1, r2)
			ret--
		}
	}
	return ret
}
