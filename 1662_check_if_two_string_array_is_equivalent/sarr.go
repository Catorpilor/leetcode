package sarr

import "strings"

func isEqual(word1, word2 []string) bool {
	// return useStd(word1, word2)
	return useIter(word1, word2)
}

func useStd(word1, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func useIter(word1, word2 []string) bool {
	n1, n2 := len(word1), len(word2)
	var i, j int
	var idx1, idx2 int
	for i < n1 && j < n2 {
		if word1[i][idx1] != word2[j][idx2] {
			return false
		}
		p1, p2 := len(word1[i]), len(word2[j])
		idx1, idx2 = idx1+1, idx2+1
		if idx1 == p1 {
			i++
			idx1 = 0
		}
		if idx2 == p2 {
			j++
			idx2 = 0
		}
	}
	return true
}
