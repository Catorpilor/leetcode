package buddy

func buddyString(a, b string) bool {
	return useStraight(a, b)
}

// time complexity O(N), space complexity O(1)
func useStraight(A, B string) bool {
	na, nb := len(A), len(B)
	if na != nb {
		return false
	}
	seen := make(map[byte]bool, 26)
	for i := range A {
		seen[A[i]] = true
	}
	if A == B && len(seen) < na {
		return true
	}
	diff := make([]int, 0, na)
	for i := 0; i < na; i++ {
		if A[i] != B[i] {
			diff = append(diff, i)
		}
	}
	return len(diff) == 2 && (A[diff[0]] == B[diff[1]] && A[diff[1]] == B[diff[0]])
}
