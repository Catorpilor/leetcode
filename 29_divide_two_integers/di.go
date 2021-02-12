package di

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func divide(dividend, divisor int) int {
	return useBit(dividend, divisor)
}

// useBit time complexity O(logN), space complexity O(1)
func useBit(dividend, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	a, b := utils.Abs(dividend), utils.Abs(divisor)
	var res, x int
	for a-b >= 0 {
		for x = 0; a-(b<<x<<1) >= 0; {
			x++
		}
		res += 1 << x
		a -= b << x
	}
	if dividend > 0 && divisor > 0 {
		return res
	}
	return -res
}
