package freq

import (
	"sort"
	"strings"
)

func freqSort(s string) string {
	return useHashmap(s)
}

func useHashmap(s string) string {
	n := len(s)
	set := make(map[byte]int, n)
	for i := range s {
		set[s[i]]++
	}
	type store struct {
		key  byte
		freq int
	}
	base := make([]store, 0, len(set))
	for k, v := range set {
		base = append(base, store{key: k, freq: v})
	}
	sort.Slice(base, func(i, j int) bool {
		return base[i].freq > base[j].freq
	})
	var sb strings.Builder
	for _, ele := range base {
		sb.WriteString(strings.Repeat(string(ele.key), ele.freq))
	}
	return sb.String()
}
