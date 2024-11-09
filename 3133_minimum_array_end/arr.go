package marr

// minArrEnd time complexity O(N), space complexity O(1)
// using bitmask to get the minimum value
func minArrEnd(n, x int) int {
	res := x
	for i := 1; i < n; i++ {
		res = (res + 1) | x
	}
	return res
}
