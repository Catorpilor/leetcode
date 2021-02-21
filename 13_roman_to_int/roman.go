package roman

func romanToInt(s string) int {
	return useSF(s)
}

// useSF time complexity O(N), space complexity O(1)
func useSF(s string) int {
	n := len(s)
	if n < 1 {
		return 0
	}
	set := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans := set[s[n-1]]
	for i := n - 2; i >= 0; i-- {
		if set[s[i]] < set[s[i+1]] {
			ans -= set[s[i]]
		} else {
			ans += set[s[i]]
		}
	}
	return ans
}
