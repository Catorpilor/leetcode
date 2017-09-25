package palin

func ValidPalindrome(s string) bool {
	if len(s) <= 2 {
		return true
	}
	if len(s) == 3 {
		if s[0] == s[1] || s[1] == s[2] || s[0] == s[2] {
			return true
		}
		return false
	}
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isPalindrome(s, l, r-1) || isPalindrome(s, l+1, r)
		}
		l, r = l+1, r-1
	}
	return true

	// odd := true
	// if n%2 == 0 {
	// 	odd = false
	// }
	// mid := n / 2

	// pre, post := s[:mid], s[mid:]
	// if odd {
	// 	post = s[mid+1:]
	// }
	// delCount := 0
	// i, j := 0, mid-1
	// for i < mid && j >= 0 {
	// 	if pre[i] != post[j] {
	// 		delCount += 1
	// 		if delCount > 1 {
	// 			return false
	// 		}
	// 	}
	// 	i, j = i+1, j-1
	// }
	// return delCount <= 1
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l, r = l+1, r-1
	}
	return true
}
