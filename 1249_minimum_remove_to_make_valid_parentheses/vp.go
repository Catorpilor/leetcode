package vp

import "strings"

func minRemove(s string) string {
	return useStack(s)
}

// useStack time complexity O(N), space complexity O(N)
func useStack(s string) string {
	sb := []byte(s)
	n := len(sb)
	st := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if sb[i] == '(' {
			st = append(st, i)
		} else if sb[i] == ')' {
			nst := len(st)
			if nst > 0 {
				// pop
				st = st[:nst-1]
			} else {
				// replace ) with *
				sb[i] = '*'
			}
		}
	}
	nst := len(st)
	for nst > 0 {
		top := st[nst-1]
		nst--
		st = st[:nst]
		sb[top] = '*'
	}
	ns := string(sb)
	return strings.ReplaceAll(ns, "*", "")
}
