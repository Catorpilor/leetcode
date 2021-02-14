package anagram

func isAnagram(s, t string) bool {
	return useHashmap(s, t)
}

// useHashmap time complexity O(N), space complexity O(N)
func useHashmap(s, t string) bool {
	n1, n2 := len(s), len(t)
	if n1 != n2 {
		return false
	}

	set := make(map[byte]int, n1)
	for i := range s {
		set[s[i]]++
	}
	for i := range t {
		if v, exists := set[t[i]]; !exists {
			return false
		} else {
			v--
			if v < 0 {
				return false
			}
			set[t[i]] = v
		}
	}
	return true
}
