package array

func minOps(n int) int {
	return useTwoPointers(n)
}

// useTwoPointers time complexity O(N), space complexity O(1)
func useTwoPointers(n int) int {
	// the array is arithmetic
	avg := n
	l, r := 0, n-1
	var ans int
	for l < r {
		ans += 2*r + 1 - avg
		l++
		r--
	}
	return ans
}
