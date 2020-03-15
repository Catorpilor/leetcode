package interleave

func isInterLeaving(s1, s2, s3 string) bool {
	return useRecursion(s1, s2, s3, 0, 0, 0)
}

// useRecursion time complexity O(2^(M+N)) space complexity O(M+N)
func useRecursion(s1, s2, s3 string, p1, p2, p3 int) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1 + n2 != n3 {
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
