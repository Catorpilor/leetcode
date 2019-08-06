package ga

import (
	"sort"
	"strings"
)

func GroupAnagrams(input []string) [][]string {
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
