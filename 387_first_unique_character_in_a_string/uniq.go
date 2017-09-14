package uniq

func FirstUniq(s string) int {
	hm := make(map[rune]int)
	for _, c := range s {
		if _, ok := hm[c]; ok {
			hm[c] += 1
		} else {
			hm[c] = 1
		}
	}
	ret := -1
	for i, v := range s {
		if hm[v] == 1 {
			ret = i
			break
		}
	}
	return ret
}
