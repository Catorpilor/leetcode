package pattern

import "strings"

func WordPattern(pattern, str string) bool {
	// split str
	cstr := strings.Split(str, " ")
	if len(cstr) != len(pattern) {
		return false
	}
	h := [128]int{}
	hm := make(map[string]struct {
		c     rune
		count int
	})
	for i, c := range pattern {
		h[c-'a'] += 1
		if val, ok := hm[cstr[i]]; ok {
			val.count += 1
			if val.c != c || val.count != h[c-'a'] {
				return false
			}
			hm[cstr[i]] = val
		} else {
			if h[c-'a'] > 1 {
				return false
			}
			hm[cstr[i]] = struct {
				c     rune
				count int
			}{
				c:     c,
				count: 1,
			}
		}
	}
	return true
}
