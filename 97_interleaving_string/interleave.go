package interleave

func isInterLeaving(s1, s2, s3 string) bool {
	// return useRecursion(s1, s2, s3, 0, 0, 0)
	return useDP(s1, s2, s3)
}

// useRecursion time complexity O(2^(M+N)) space complexity O(M+N)
func useRecursion(s1, s2, s3 string, p1, p2, p3 int) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}
	if p1 == n1 && p2 == n2 && p3 == n3 {
		return true
	}
	// if p3 == n3 {
	// 	return false
	// }
	var ret bool
	if p1 < n1 && s1[p1] == s3[p3] {
		ret = useRecursion(s1, s2, s3, p1+1, p2, p3+1) || ret
	}
	if p2 < n2 && s2[p2] == s3[p3] {
		ret = useRecursion(s1, s2, s3, p1, p2+1, p3+1) || ret
	}
	return ret
}

// useDP time complexity O(MN), space complexity O(MN)
func useDP(s1, s2, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}
	// dp[i][j] represents if  s1[:i], s2[:j],  s3[:i+j] are interleaving strings.
	dp := make([][]bool, n1+1)
	for i := range dp {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true // both empty
	for j := 1; j <= n2; j++ {
		if s2[j-1] == s3[j-1] {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i <= n1; i++ {
		if s1[i-1] == s3[i-1] {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s3[i+j-1] == s1[i-1] {
				dp[i][j] = dp[i-1][j] || dp[i][j]
			}
			if s3[i+j-1] == s2[j-1] {
				dp[i][j] = dp[i][j-1] || dp[i][j]
			}
		}
	}
	return dp[n1][n2]
}
