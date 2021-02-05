package path

import (
	"bytes"
	"strings"
)

func simplifyPath(path string) string {
	return useStack(path)
}

// useStack time complexity O(N), space complexity O(N)
func useStack(path string) string {
	path = strings.Replace(path, "//", "/", -1)
	dirs := strings.FieldsFunc(path, func(c rune) bool {
		return c == '/'
	})
	st := make([]string, 0, len(dirs))
	var sb bytes.Buffer
	sb.WriteByte('/')
	for _, dir := range dirs {
		if dir == ".." {
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
		} else if dir == "." {
			continue
		} else {
			st = append(st, dir)
		}
	}
	for i, dir := range st {
		sb.WriteString(dir)
		if i < len(st)-1 {
			sb.WriteByte('/')
		}
	}
	return sb.String()
}
