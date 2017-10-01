package lp

func LongestPalindrome(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	count := 0
	hm := make(map[rune]int)
	for _, c := range s {
		if _, ok := hm[c]; ok {
			count += 1
			delete(hm, c)
		} else {
			hm[c] = 1
		}
	}
	if len(hm) != 0 {
		return count*2 + 1
	}
	return count * 2
}

func LongestPalindrome2(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	h := [58]int{}
	for _, c := range s {
		h[c-'A'] += 1
	}
	eve, odd := 0, 0
	for _, v := range h {
		// this is the trick part
		// but better than my first version
		// which i thought
		// if it is odd plus v - 1
		// otherwise plus v
		eve += v & ^1
		if v%2 == 1 {
			// only add once
			odd = 1
		}
	}
	return eve + odd
}
