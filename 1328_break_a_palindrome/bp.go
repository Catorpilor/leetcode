package bp

func breakPalindrome(s string) string {
	return iteratorV2(s)
}

// bruteForce time complexity O(N^2)
func bruteForce(s string) string {
	n := len(s)
	var res string
	if n == 1 {
		return res
	}
	for i := range s {
		for c := 'a'; c <= 'z'; c++ {
			sb := []byte(s)
			if sb[i] != byte(c) {
				sb[i] = byte(c)
				tmpsb := string(sb)
				if !isPalindrome(tmpsb) {
					if res == "" {
						res = tmpsb
					} else {
						if tmpsb < res {
							res = tmpsb
						}
					}
				}
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// iterator time complexity is O(N)
func iterator(s string) string {
	n := len(s)
	if n <= 1 {
		return ""
	}
	sb := []byte(s)
	for i := 0; i < n; i++ {
		j := n - 1 - i
		if i == j {
			// n is odd, skip this character
			// for example aacaa, if we change c to a, this is still a palindrome
			continue
		}
		if sb[i] != byte('a') {
			sb[i] = byte('a')
			return string(sb)
		}
	}
	// all 'a's
	sb[n-1] = byte('b')
	return string(sb)
}

// iteratorV2 time complexity is O(N)
func iteratorV2(s string) string {
	n := len(s)
	if n <= 1 {
		return ""
	}
	sb := []byte(s)
	// s is a palindrome, we just check half the string
	for i := 0; i < n/2; i++ {
		if sb[i] != byte('a') {
			sb[i] = byte('a')
			return string(sb)
		}
	}
	// all 'a's
	sb[n-1] = byte('b')
	return string(sb)
}
