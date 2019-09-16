package pis

func checkInclusion(s1, s2 string) bool {
	// return hashmap(s1, s2)
	// return array(s1, s2)
	return slideWindow(s1, s2)
}

func slideWindow(s1, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 == 0 || n2 == 0 || n1 > n2 {
		return false
	}
	c1, c2 := make([]int, 26), make([]int, 26)
	for i := range s1 {
		c1[int(s1[i]-'a')]++
		c2[int(s2[i]-'a')]++
	}
	for i := 0; i < n2-n1; i++ {
		if compArray(c1, c2) {
			return true
		}
		c2[int(s2[i+n1]-'a')]++
		c2[int(s2[i]-'a')]--
	}
	return compArray(c1, c2)
}

func array(s1, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 == 0 || n2 == 0 || n1 > n2 {
		return false
	}
	cache1 := make([]int, 26)
	for i := range s1 {
		cache1[int(s1[i]-'a')]++
	}
	for i := 0; i <= n2-n1; i++ {
		cache2 := make([]int, 26)
		for j := i; j < i+n1; j++ {
			cache2[int(s2[j]-'a')]++
		}
		if compArray(cache1, cache2) {
			return true
		}
	}
	return false
}

func compArray(a1, a2 []int) bool {
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func hashmap(s1, s2 string) bool {
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
