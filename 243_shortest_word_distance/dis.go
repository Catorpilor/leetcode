package distance

func ShortestDistance(words []string, w, s string) int {
	idxw, idxs, ret := -1, -1, len(words)
	for i, k := range words {
		if k == w {
			idxw = i
		}
		if k == s {
			idxs = i
		}
		if idxs != -1 && idxw != -1 {
			ret = min(ret, abs(idxw-idxs))
		}
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
