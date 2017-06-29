package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefix := ""
	for idx := 0; len(strs) > 0; idx++ {
		for i := 0; i < len(strs); i++ {
			if idx >= len(strs[i]) || (i > 0 && strs[i][idx] != strs[i-1][idx]) {
				return prefix
			}
		}
		prefix += string(strs[0][idx])
	}
	return prefix

}

func main() {
	fmt.Println("vim-go")
	fmt.Println(longestCommonPrefix([]string{"a", "ab", "ac"}))
}
