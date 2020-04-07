package count

func countElements(arr []int) int {
	return useHashmap(arr)
}

func useHashmap(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}
	return n
}
