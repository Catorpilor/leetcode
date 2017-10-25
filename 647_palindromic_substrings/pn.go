package pn

func CountSubStrings(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var count int
	for i := 0; i < len(s); i++ {
		checkPanlidrome(s, i, i, &count)
		checkPanlidrome(s, i, i+1, &count)
	}
	return count
}

func checkPanlidrome(s string, l, r int, count *int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		*count, l, r = *count+1, l-1, r+1
	}
}
