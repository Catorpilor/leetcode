package prefix

import "strings"

func isPrefix(s, word string) int {
	return useOnePass(s, word)
}

// useOnePass time complexity O(NM), space complexity O(1)
// N represents the number of strings. M means the length of word
func useOnePass(str, word string) int {
	ss := strings.Fields(str)
	res := -1
	for i := range ss { // O(N)
		if strings.HasPrefix(ss[i], word) { // O(M)
			res = i + 1
			break
		}
	}
	return res
}
