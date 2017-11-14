package ascii

import (
	"github.com/catorpilor/leetcode/utils"
)

func MinDelAscii(s1, s2 string) int {
	n1, n2 := len(s1), len(s2)
	if n1 == 0 && n2 == 0 {
		return 0
	}
	// dp[i][j] represent the minimal delete for s1.substr(0,i) and s2.substr(0,j)
	// similar to get longest common sequence
	// but we store the cost instead
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	for j := 1; j <= n2; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}
	for i := 1; i <= n1; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
		for j := 1; j <= n2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = utils.Min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}
	return dp[n1][n2]
}

func MinDelAscii2(s1, s2 string) int {
	n1, n2 := len(s1), len(s2)
	if n1 == 0 && n2 == 0 {
		return 0
	}
	dp := make([]int, n2+1)
	for i := 1; i <= n2; i++ {
		dp[i] = dp[i-1] + int(s2[i-1])
	}
	var temp1, temp2 int
	for i := 1; i <= n1; i++ {
		temp1 = dp[0]
		dp[0] += int(s1[i-1])
		for j := 1; j <= n2; j++ {
			temp2 = dp[j]
			if s1[i-1] == s2[j-1] {
				dp[j] = temp1
			} else {
				dp[j] = utils.Min(dp[j]+int(s1[i-1]), dp[j-1]+int(s2[j-1]))
			}
			temp1 = temp2
		}
	}
	return dp[n2]
}
