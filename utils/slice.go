package utils

// DeleteFromIntSlice delete target from slice a
func DeleteFromIntSlice(a []int, target int) []int {
	n := len(a)
	for i, v := range a {
		if v == target {
			copy(a[i:], a[i+1:])
			a = a[:n-1]
			return a
		}
	}
	return a
}
