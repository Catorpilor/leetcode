package bsc

import "bytes"

func backspaceCompare(s, t string) bool {
	return useStack(s, t)
}

// useStack time complexity O(M+N), space complexity O(M+N)
func useStack(s, t string) bool {
	ns, nt := len(s), len(t)
	st1, st2 := make([]byte, 0, ns), make([]byte, 0, nt)
	var nst1, nst2 int
	for i := range s {
		nst1 = len(st1)
		if s[i] == '#' {
			if nst1 > 0 {
				st1 = st1[:nst1-1]
			}
			continue
		}
		st1 = append(st1, s[i])
	}
	for j := range t {
		nst2 = len(st2)
		if t[j] == '#' {
			if nst2 > 0 {
				st2 = st2[:nst2-1]
			}
			continue
		}
		st2 = append(st2, t[j])
	}
	return bytes.Equal(st1, st2)
}
