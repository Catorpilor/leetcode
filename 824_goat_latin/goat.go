package goat

import (
	"bytes"
	"strings"
)

func toGoatLatin(s string) string {
	return useStraight(s)
}

func useStraight(s string) string {
	sb := []byte(s)
	sbs := bytes.Fields(sb)
	var bb bytes.Buffer
	for i, w := range sbs {
		if isVowel(w[0]) {
			bb.Write(w)
		} else {
			w = append(w, w[0])
			bb.Write(w[1:])
		}
		bb.WriteString("ma" + strings.Repeat("a", i+1))
		if i < len(sbs)-1 {
			bb.WriteString(" ")
		}
	}
	return bb.String()
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'o', 'i', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
