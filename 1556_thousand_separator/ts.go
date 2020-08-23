package ts

import (
	"bytes"
	"strconv"
)

func addSeps(num int) string {
	return useStraight(num)
}

// useStraight time complexity O(N), space complexity O(N)
func useStraight(num int) string {
	s := strconv.Itoa(num)
	var tmp []string
	ns := len(s)
	end := ns
	for ns >= 3 {
		tmp = append(tmp, s[ns-3:end])
		end = ns - 3
		ns -= 3
	}
	if ns > 0 {
		tmp = append(tmp, s[:ns])
	}
	var sb bytes.Buffer
	for k := len(tmp) - 1; k >= 0; k-- {
		sb.WriteString(tmp[k])
		if k > 0 {
			sb.WriteByte('.')
		}
	}
	return sb.String()
}
