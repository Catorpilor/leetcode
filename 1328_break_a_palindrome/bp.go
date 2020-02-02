package bp

func breakPalindrome(s string) string {
	return bruteForce(s)
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
