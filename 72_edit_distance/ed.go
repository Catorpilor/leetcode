package ed

import "github.com/catorpilor/LeetCode/utils"

func minDistance(word1, word2 string) int {
	n1, n2 := len(word1), len(word2)
	if n1 == 0 {
		return n2
	}
	if n2 == 0 {
		return n1
	}
	// edge cases when word1 and word2 are equal
	if word1 == word2 {
		return 0
	}
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	// init row 0
	// for example word1="", iterator word2, update dp[0][j] to j
	for j := 1; j <= n2; j++ {
		dp[0][j] = j
	}
	// init column 0
	// for example word2="", iterator word1, update dp[i][0] to i
	for i := 1; i <= n1; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				// at current pos i, j word1[i-1] == word2[j-1]
				// no need to change
				dp[i][j] = dp[i-1][j-1]
			} else {
				//dp[i-1][j-1]+1 ----> the replace op
				// dp[i][j-1]+1, dp[i-1][j]+1  ---> delete & insert op
				// udpate dp[i][j] to the minimal of  three steps
				dp[i][j] = utils.Min(dp[i-1][j-1]+1, utils.Min(dp[i-1][j]+1, dp[i][j-1]+1))
			}
		}
	}

	return dp[n1][n2]
}
