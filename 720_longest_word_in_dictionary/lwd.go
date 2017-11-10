package lwd

import "sort"

func LongestWord(words []string) string {
	n := len(words)
	if n < 1 {
		return ""
	}
	if n == 1 {
		return words[0]
	}
	// sort words
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})
	var ret string
	hash := make(map[string]bool)
	for _, w := range words {
		nw := len(w)
		if nw == 1 || hash[w[:nw-1]] {
			if nw > len(ret) {
				ret = w
			}
			hash[w] = true
		}
	}
	return ret
}
