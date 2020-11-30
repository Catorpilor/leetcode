package flipping

func flipAndInvert(A [][]int) [][]int {
	return useBruteForce(A)
}

// useBruteForce time complexity O(MN), space complexity O(1)
func useBruteForce(A [][]int) [][]int {
	m := len(A)
	if m < 1 {
		return A
	}
	n := len(A[0])
	if n < 1 {
		return A
	}
	for i := 0; i < m; i++ {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			A[i][l], A[i][r] = A[i][r], A[i][l]
		}
		for j := range A[i] {
			A[i][j] ^= 1
		}
	}
	return A
}
