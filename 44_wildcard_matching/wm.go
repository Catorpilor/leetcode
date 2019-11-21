package wm

func isMatchTLE(s, p string) bool {
	ns, np := len(s), len(p)
	// empty case
	if ns == 0 {
		return np == 0 || allAsterisk(p)
	}
	if np == 0 {
		return false
	}
	if allAsterisk(p) {
		return true
	}
	if p[0] == '*' {
		var ret bool
		for i := range s {
			ret = ret || isMatch(s[i:], p[1:])
			if ret {
				break
			}
		}
		return ret
	} else if p[0] == s[0] || p[0] == '?' {
		return isMatch(s[1:], p[1:])
	}

	return false
}

func allAsterisk(p string) bool {
	if len(p) == 0 {
		return false
	}
	for _, c := range p {
		if c != '*' {
			return false
		}
	}
	return true
}

func isMatch(s, p string) bool {
	var i, j int
	ast_pos, pos := -1, 0
	if len(p) < 1 {
		return false
	}

	for i < len(s) {
		if s[i] == p[j] || p[j] == '?' {
			i++
			j++
			continue
		}
		if p[j] == '*' {
			ast_pos = j
			j++
			pos = i
			continue
		}
		if ast_pos != -1 {
			j = ast_pos + 1
			pos++
			i = pos
			continue
		}
		return false
	}
	for j < len(p) && p[j] == '*' {
		j++
	}
	return j == len(p)
}
