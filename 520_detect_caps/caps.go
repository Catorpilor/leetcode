package caps

func detectCaps(word string) bool {
	return useOnepass(word)
}

// useOnepass time complexity O(N), space complexity O(1)
func useOnepass(word string) bool {
	seenNonCaps, allCaps := false, false
	for i, c := range word {
		if c >= 'A' && c <= 'Z' {
			if i > 0 && seenNonCaps {
				return false
			}
			if i > 0 {
				allCaps = true
			}
		} else {
			seenNonCaps = true
			if allCaps {
				return false
			}
		}
	}
	return true
}
