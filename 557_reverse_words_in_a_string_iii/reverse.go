package reverse

import (
	"bytes"
	"strings"
)

func ReverseWords(s string) string {
	// brute force
	strs := strings.Split(s, " ")
	var ret bytes.Buffer
	for i, ss := range strs {
		// reverse each word
		ret.WriteString(reverse(ss))
		if i != len(strs)-1 {
			ret.WriteRune(' ')
		}
	}
	return ret.String()
}

func reverse(s string) string {
	bs := []rune(s)
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}
