package lp

func LongestPalindrome(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	hm := make(map[rune]int)
	for _, c := range s {
		hm[c] += 1
	}
	ret := 0
	added := false
	for _, v := range hm {
		if v%2 != 0 {
			ret += v - 1
			if v == 1 && !added {
				ret += 1
				added = true
			}
		} else {
			ret += v
		}
	}
	return ret
}
