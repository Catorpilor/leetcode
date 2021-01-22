package ds

import (
	"sort"

	"github.com/google/go-cmp/cmp"
)

func isClose(word1, word2 string) bool {
	return useHashMap(word1, word2)
}

// useHashMap time complexity O(N), space complexity O(N)
func useHashMap(word1, word2 string) bool {
	n1, n2 := len(word1), len(word2)
	if n1 != n2 {
		return false
	}
	if n1 == 0 {
		return true
	}
	occ := make(map[byte]int, n1)
	for i := range word1 {
		occ[word1[i]]++
	}
	occ2 := make(map[byte]int, n2)
	for i := range word2 {
		if occ[word2[i]] == 0 {
			return false
		}
		occ2[word2[i]]++
	}
	val1, val2 := make([]int, 0, n1), make([]int, 0, n2)
	for _, v := range occ {
		val1 = append(val1, v)
	}
	for _, v := range occ2 {
		val2 = append(val2, v)
	}
	sort.Ints(val1)
	sort.Ints(val2)
	if diff := cmp.Diff(val1, val2); diff != "" {
		return false
	}
	return true
}
