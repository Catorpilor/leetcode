package islands

func NumberOfIslands(m, n int, pos [][]int) []int {
	if m*n == 0 || len(pos) == 0 {
		return nil
	}
	ids := make([]int, m*n)
	wz := make([]int, m*n)
	ret := make([]int, 0, len(pos))
	hm := make(map[int]bool)
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
	for i := range pos {
		t := pos[i][0]*n + pos[i][1]
		ids[t] = t
		wz[t] = 1
		// hm[t] = true
		ret = append(ret, i+1)
	}
	hm[pos[0][0]*n+pos[0][1]] = true
	for k := 1; k < len(pos); k++ {
		ret[k] = ret[k-1] + 1
		i, j := pos[k][0], pos[k][1]
		hm[i*n+j] = true
		if i > 0 && hm[(i-1)*n+j] {
			if union(i*n+j, (i-1)*n+j) {
				ret[k]--
			}
		}
		if j > 0 && hm[i*n+j-1] {
			if union(i*n+j, i*n+j-1) {
				ret[k]--
			}
		}
		if i+1 < m && hm[(i+1)*n+j] {
			if union(i*n+j, (i+1)*n+j) {
				ret[k]--
			}
		}
		if j+1 < n && hm[i*n+j+1] {
			if union(i*n+j, i*n+j+1) {
				ret[k]--
			}
		}
	}
	return ret
}
