package pis

func checkInclusion(s1, s2 string) bool {
	// time complexity O((n2-n1)*n1+n1)
	// space complexity O((n2-n1)*n1)
	n1, n2 := len(s1), len(s2)
	if n1 == 0 || n2 == 0 || n1 > n2 {
		return false
	}
	cache1 := make(map[byte]int, n1)
	for i := range s1 {
		cache1[s1[i]]++
	}
	for i := 0; i <= n2-n1; i++ {
		cache2 := make(map[byte]int, n1)
		for j := i; j < i+n1; j++ {
			cache2[s2[j]]++
		}
		if comp(cache1, cache2) {
			return true
		}
	}
	return false
}

func comp(m1, m2 map[byte]int) bool {
	for k := range m1 {
		if m1[k]-m2[k] != 0 {
			return false
		}
	}
	return true
}
