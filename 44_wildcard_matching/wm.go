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
			ret = ret || isMatchTLE(s[i:], p[1:])
			if ret {
				break
			}
		}
		return ret
	} else if p[0] == s[0] || p[0] == '?' {
		return isMatchTLE(s[1:], p[1:])
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
	// if len(p) < 1 {
	// 	return false
	// }

	for i < len(s) {
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < len(p) && p[j] == '*' {
			// the zero case
			ast_pos = j
			j++
			pos = i
		} else if ast_pos != -1 {
			// s[i] != p[j] and there is an asterisk exists.
			j = ast_pos + 1
			// one or more case when there is a '*'
			// move forward i
			pos++
			i = pos
		} else {
			// s[i] != p[j] && no '*' found
			// just return false
			return false
		}
	}
	for j < len(p) && p[j] == '*' {
		j++
	}
	return j == len(p)
}

func dp(s, p string) bool {
	ns, np := len(s), len(p)
	if p == s || p == "*" {
		return true
	}
	if ns == 0 || np == 0 {
		return false
	}
	dpm := make([][]bool, np+1)
	for i := range dpm {
		dpm[i] = make([]bool, ns+1)
	}
	dpm[0][0] = true
	for pIdx := 1; pIdx < np+1; pIdx++ {
		// p[pIdx-1] == '*'
		if p[pIdx-1] == '*' {
			sIdx := 1
			// dpm[pIdx-1][sIdx-1] is the previous match, one character before
			// here we find the first sIdx in `s` with previous match.
			for !dpm[pIdx-1][sIdx-1] && sIdx < ns+1 {
				sIdx++
			}
			// if s match p
			// then s match p* as well
			dpm[pIdx][sIdx-1] = dpm[pIdx-1][sIdx-1]
			// if s match p
			// then s(other characters) matches p* as well
			for sIdx < ns+1 {
				dpm[pIdx][sIdx] = true
				sIdx++
			}
		} else if p[pIdx-1] == '?' {
			for sIdx := 1; sIdx < ns+1; sIdx++ {
				dpm[pIdx][sIdx] = dpm[pIdx-1][sIdx-1]
			}
		} else {
			for sIdx := 1; sIdx < ns+1; sIdx++ {
				// match is possible if there is a previous match and
				// current character is same
				dpm[pIdx][sIdx] = dpm[pIdx-1][sIdx-1] && p[pIdx-1] == s[sIdx-1]
			}
		}
	}
	return dpm[np][ns]
}
