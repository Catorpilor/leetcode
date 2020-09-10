package ways

func numOfWays(s string) int {
	return useMath(s)
}

// useMath time complexity O(N), space complexity O(N)
func useMath(s string) int {
	n := len(s)
	ones := make([]int, 0, n)
	for i := range s {
		if s[i] == '1' {
			ones = append(ones, i)
		}
	}
	mod := int(1e9 + 7)
	// for a length = n string, there are n-1 slots, so the first cut can have (n-1) choices,
	// and the second cut can have (n-2) choices, but there are duplicates
	if len(ones) == 0 {
		return ((n - 1) * (n - 2) / 2) % mod
	}
	if len(ones)%3 != 0 {
		return 0
	}
	target := len(ones) / 3
	return ((ones[target] - ones[target-1]) * (ones[target*2] - ones[target*2-1])) % mod
}
