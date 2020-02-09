package seq

import (
	"fmt"
	"math"
	"strconv"

	"github.com/catorpilor/leetcode/utils"
)

func splitIntoFib(s string) []int {
	var res []int
	n := len(s)
	if n < 3 {
		return res
	}
	// same idea as additive number
	// i, j represents the length of the x1, x2
	for i := 1; i <= n/2; i++ {
		for j := 1; utils.Max(j, i) <= n-i-j; j++ {
			r, f := isValid(i, j, n, s)
			if f {
				return r
			}
		}
	}
	return res
}

func isValid(i, j, n int, s string) (res []int, f bool) {
	if s[0] == '0' && i > 1 {
		return
	}
	if s[i] == '0' && j > 1 {
		return
	}
	x1, _ := strconv.ParseInt(s[:i], 10, 64)
	x2, _ := strconv.ParseInt(s[i:i+j], 10, 64)
	tmp1, tmp2 := x1, x2
	var sum string
	for start := i + j; start != n; start += len(sum) {
		tmp2 += tmp1
		tmp1 = tmp2 - tmp1
		if tmp1 > math.MaxInt32 || tmp2 > math.MaxInt32 {
			return
		}
		sum = strconv.FormatInt(tmp2, 10)
		tn := start + len(sum)
		if tn > n {
			tn = n
		}
		if sum != s[start:tn] {
			return
		}
		fmt.Printf("x1:%d, x2:%d,sum:%s,res:%v, start: %d, tn: %d\n", x1, x2, sum, res, start, tn)
		// res = append(res, x1, x2, tmp2)
		if len(res) == 0 {
			res = append(res, int(x1), int(x2), int(tmp2))
		} else {
			res = append(res, int(tmp2))
		}
		// if tn == n {
		// 	res = append(res, tmp2)
		// }
		x1 = tmp1
		x2 = tmp2
	}
	f = true
	return
}
