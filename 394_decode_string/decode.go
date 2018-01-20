package decode

import (
	"bytes"

	"github.com/catorpilor/leetcode/utils"
)

func DecodeString(s string) string {
	st, res := utils.NewStack(), utils.NewStack()
	var b bytes.Buffer
	var idx, count int
	for idx < len(s) {
		if isDigit(s[idx]) {
			// if it is a digit
			count = 0
			for idx < len(s) && isDigit(s[idx]) {
				count = 10*count + int(s[idx]-'0')
				idx++
			}
			st.Push(count)
		} else if s[idx] == '[' {
			temp := make([]byte, b.Len())
			copy(temp, b.Bytes())
			res.Push(temp)
			b.Reset()
			idx++
		} else if s[idx] == ']' {
			t := res.Pop().([]byte)
			cnt := st.Pop().(int)
			for i := 0; i < cnt; i++ {
				t = append(t, b.Bytes()...)
			}
			b.Reset()
			b.Write(t)
			idx++
		} else {
			b.WriteByte(s[idx])
			idx++
		}
	}
	return b.String()
}

func isDigit(b byte) bool {
	if int(b-'0') >= 0 && int(b-'0') <= 9 {
		return true
	}
	return false
}
