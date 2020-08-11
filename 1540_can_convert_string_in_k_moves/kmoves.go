package kmoves

func canConvert(s, t string, k int) bool {
	return useOnePass(s, t, k)
}

// useOnePass time complexity O(N), space complexity O(1)
func useOnePass(s, t string, k int) bool {
	ns, nt := len(s), len(t)
	if ns != nt {
		return false
	}

	base := make([]int, 26)
	for j := range s {
		if s[j] == t[j] {
			continue
		}
		diff := int(t[j]-'a') - int(s[j]-'a')
		if diff < 0 {
			diff += 26
		}
		// fmt.Printf("diff: %d, t[j]:%c, s[j]:%c, base[diff]:%d\n", diff, t[j], s[j], base[diff])
		if diff+base[diff] > k {
			return false
		}
		// the next round we have to do one more rotating. +26
		base[diff] += 26

	}

	return true
}
