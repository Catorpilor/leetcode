package pl

func labels(s string) []int {
	return useGreedy(s)
}

// useGreedy time complexity O(N), space complexity O(1)
func useGreedy(s string) []int {
	var ans []int
	n := len(s)
	if n < 1 {
		return ans
	}
	set := make([]int, 26)
	for i := range s {
		idx := int(s[i] - 'a')
		set[idx] = i
	}
	start := 0
	end := set[int(s[0]-'a')]
	for i := 1; i < n; i++ {
		idx := int(s[i] - 'a')
		if set[idx] > end {
			if i > end {
				ans = append(ans, end-start+1)
				start = i
			}
			end = set[idx]
		}
	}
	ans = append(ans, n-start)
	return ans
}
