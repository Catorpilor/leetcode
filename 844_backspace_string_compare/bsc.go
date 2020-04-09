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

// useTwoPointers time complexity O(M+N), space complexity O(1)
func useTwoPointers(s, t string) bool {
	ns, nt := len(s), len(t)
	i, j := ns-1, nt-1
	skipS, skipT := 0, 0
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 && s[i] != t[j] {
			return false
		}
		if (i >= 0) != (j >= 0) {
			return false
		}
		i--
		j--
	}
	return true
}
