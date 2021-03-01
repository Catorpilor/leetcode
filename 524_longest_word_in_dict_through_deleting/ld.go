package ld

func longestWord(s string, dict []string) string {
	return useTwoPointers(s, dict)
}

// useTwoPointers time complexity O(MN), space complexity O(1)
func useTwoPointers(s string, dict []string) string {
	m := len(s)
	var ans string
	for i := range dict {
		cur := dict[i]
		var si, ci int
		for si < len(cur) && ci < m {
			if cur[si] == s[ci] {
				si++
			}
			ci++
		}
		if si == len(cur) {
			// find one in s
			if len(ans) < len(cur) || len(ans) == len(cur) && cur < ans {
				ans = cur
			}
		}
	}
	return ans
}
