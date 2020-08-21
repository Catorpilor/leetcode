package array

func sortArrayByParity(arr []int) []int {
	return useOnePass(arr)
}

// useOnePass time complexity O(N), space complexity O(1)
func useOnePass(arr []int) []int {
	n := len(arr)
	l, r := 0, n-1
	for l < r {
		if arr[l]%2 == 0 {
			l++
		} else {
			arr[l], arr[r] = arr[r], arr[l]
			// arr[r] is odd now
			r--
		}
	}
	return arr
}
