package validp

func isValid(s string) bool {
	return useStack(s)
}

// useStack time complexity O(N), space complexity O(N)
func useStack(s string) bool {
	n := len(s)
	st := make([]rune, 0, n)
	for i := range s {
		switch s[i] {
		case '(', '{', '[':
			st = append(st, s[i])
		case ')':
			nst := len(st)
			if nst == 0 || st[nst-1] != '(' {
				return false
			}
			st = st[:nst-1]
		case '}':
			nst := len(st)
			if nst == 0 || st[nst-1] != '{' {
				return false
			}
			st = st[:nst-1]
		case ']':
			nst := len(st)
			if nst == 0 || st[nst-1] != '[' {
				return false
			}
			st = st[:nst-1]
		default:
			return false
		}
	}
	return len(st) == 0
}
