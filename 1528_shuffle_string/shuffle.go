package shuffle

func shuffleString(s string, indices []int) string {
	return useTraverse(s, indices)
}

func useTraverse(s string, indices []int) string {
	n := len(s)
	if n < 1 {
		return s
	}
	sb := []byte(s)
	undo := make([][]int, 26)
	for i := 0; i < n; i++ {
		if i == indices[i] {
			continue
		}
		from, to := i, indices[i]
		if indices[to] == from {
			// swap
			sb[from], sb[to] = sb[to], sb[from]

		} else {
			undo[s[to]-'a'] = append(undo[s[to]-'a'], indices[to])
			sb[to] = sb[from]
		}
		indices[to] = to
	}
	for i, pp := range undo {
		for _, p := range pp {
			sb[p] = byte('a' + i)
		}
	}
	return string(sb)
}
