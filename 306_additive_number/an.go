package an

import (
	"strconv"

	"github.com/catorpilor/leetcode/utils"
)

func isAdditive(s string) bool {
	n := len(s)
	// i represent the first number's length
	// j represent the second number's length
	for i := 1; i <= n/2; i++ {
		for j := 1; utils.Max(j, i) <= n-i-j; j++ {
			if isValid(i, j, n, s) {
				return true
			}
		}
	}
	return false
}

func isValid(i, j, n int, s string) bool {
	if s[0] == '0' && i > 1 {
		return false
	}
	if s[i] == '0' && j > 1 {
		return false
	}
	x1, _ := strconv.ParseInt(s[:i], 10, 64)
	x2, _ := strconv.ParseInt(s[i:i+j], 10, 64)
	var sum string
	// for example s = 199100199
	// i = 1, j = 2
	// x1 = 1, x2 = 99
	for start := i + j; start != n; start += len(sum) {
		// this is brilliant
		x2 += x1     // x2 becomes the sum
		x1 = x2 - x1 // x1 <-> original x2
		sum = strconv.FormatInt(x2, 10)
		end := start + len(sum)
		if end > n {
			end = n
		}
		if sum != s[start:end] {
			return false
		}
	}
	return true
}
