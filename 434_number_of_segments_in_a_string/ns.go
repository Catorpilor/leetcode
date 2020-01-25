package ns

import "strings"

func countSegments(s string) int {
	f := func(r rune) bool {
		return r == ' '
	}
	return len(strings.FieldsFunc(s, f))
}
