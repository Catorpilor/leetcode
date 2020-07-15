package words

import (
	"strings"
)

func reverseWords(s string) string {
	return useStraight(s)
}

// useStraight time complexity O(N), space complexity O(N)
func useStraight(s string) string {
	n := len(s)
	if n < 1 {
		return s
	}
	segs := strings.Fields(s)
	l, r := 0, len(segs)-1
	for l < r {
		segs[l], segs[r] = segs[r], segs[l]
		l++
		r--
	}
	return strings.Join(segs, " ")
}
