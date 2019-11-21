package wm

func isMatch(s, p string) bool {
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
