package sarr

import "strings"

func isEqual(word1, word2 []string) bool {
	return useStd(word1, word2)
}

func useStd(word1, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}
