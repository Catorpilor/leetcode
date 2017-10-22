package decode

import (
	"strings"

	"github.com/catorpilor/leetcode/utils"
)

const Z = "26"

func NumEncodings(t string) int {
	n := len(t)
	if n == 0 || t[0] == '0' {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if t[i-2] == '0' && t[i-1] == '0' {
			return 0
		}
		if strings.Compare(t[i-2:i], Z) == 1 {
			if t[i-1] == '0' {
				return 0
			}
			dp[i] = utils.Max(dp[i-1], dp[i-2])
		} else {
			if t[i-2] == '0' {
				dp[i] = dp[i-1]
			} else if t[i-1] == '0' {
				dp[i] = dp[i-2]
			} else {
				dp[i] = dp[i-1] + dp[i-2]
			}
		}
	}
	return dp[n]
}

func NumEncodings2(t string) int {
	n := len(t)
	if n == 0 || t[0] == '0' {
		return 0
	}
	dp := make([]int, n+1)
	dp[n] = 1
	dp[n-1] = 1
	if t[n-1] == '0' {
		dp[n-1] = 0
	}
	for i := n - 2; i >= 0; i-- {
		if t[i] != '0' {
			if strings.Compare(t[i:i+2], Z) == 1 {
				dp[i] = dp[i+1]
			} else {
				dp[i] = dp[i+1] + dp[i+2]
			}
		}
	}
	return dp[0]
}
