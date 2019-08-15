package ga

import (
	"sort"
	"strconv"
	"strings"
)

func GroupAnagrams(input []string) [][]string {
	// return ga1(input)
	return ga2(input)
}

// time complex O(nKLogK) K is the max length of str in input
// space Complexity O(nK)
func ga1(input []string) [][]string {
	n := len(input)
	lc := make(map[string][]string, n)
	for _, str := range input {
		// old school
		// sort str alphabetically
		k := sortString(str)
		if v, exists := lc[k]; !exists {
			lc[k] = []string{str}
		} else {
			lc[k] = append(v, str)
		}
	}
	out := make([][]string, 0, len(lc))
	for _, v := range lc {
		out = append(out, v)
	}

	return out
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func ga2(str []string) [][]string {
	// use bucket to count existence as key without sorting
	n := len(str)
	lc := make(map[string][]string, n)
	for _, st := range str {
		bucket := make([]int, 26)
		for i := range st {
			bucket[st[i]-'a']++
		}
		// build bucket as key
		var sb strings.Builder
		for i := 0; i < 26; i++ {
			sb.WriteRune('#')
			sb.WriteString(strconv.FormatInt(int64(bucket[i]), 10))
		}
		key := sb.String()
		cv := lc[key]
		lc[key] = append(cv, st)
	}
	out := make([][]string, 0, len(lc))
	for _, v := range lc {
		out = append(out, v)
	}
	return out
}
