package anagram

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hm := make(map[rune]int)
	for _, c := range s {
		if _, ok := hm[c]; !ok {
			hm[c] = 1
		} else {
			hm[c] += 1
		}
	}
	for _, k := range t {
		if v, ok := hm[k]; !ok {
			return false
		} else {
			if v == 1 {
				delete(hm, k)
			} else {
				hm[k] -= 1
			}
		}
	}

	return len(hm) == 0
}

func IsAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	h := [26]int{}
	for _, c := range s {
		h[c-'a'] += 1
	}

	for _, k := range t {
		h[k-'a'] -= 1
	}

	for _, v := range h {
		if v != 0 {
			return false
		}
	}
	return true
}
