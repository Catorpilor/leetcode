package file

import (
	"fmt"
	"strconv"
)

func getFolderNames(names []string) []string {
	return useHashmap(names)
}

// useHashmap  time complexity O(N), space complexity O(N)
func useHashmap(names []string) []string {
	n := len(names)
	if n < 1 {
		return names
	}
	set := make(map[string]int, n)
	ans := make([]string, n)
	for i, name := range names {
		if v, exists := set[name]; !exists {
			ans[i] = name
			set[name] = 1
		} else {
			cur := v
			for {
				cname := fmt.Sprintf("%s(%s)", name, strconv.Itoa(cur))
				if _, exists := set[cname]; !exists {
					ans[i] = cname
					set[cname] = 1
					break
				}
				cur++
			}
			cur++
			set[name] = cur
		}
	}
	return ans
}
