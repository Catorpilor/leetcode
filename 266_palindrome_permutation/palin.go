package palin

func CanPermutePalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	hm := make(map[rune]bool)
	for _, c := range s {
		if _, ok := hm[c]; ok {
			delete(hm, c)
		} else {
			hm[c] = true
		}
	}
	return len(hm) == 0 || len(hm) == 1
}

func CanPermutePalindrome2(s string) bool {
	if len(s) <= 1 {
		return true
	}
	hm := [256]int{}
	for _, c := range s {
		hm[c-'a'] += 1
	}
	cnt := 0
	for _, c := range s {
		cnt += hm[c-'a'] % 2
	}
	return cnt <= 1
}
