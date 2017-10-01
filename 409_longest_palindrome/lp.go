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
