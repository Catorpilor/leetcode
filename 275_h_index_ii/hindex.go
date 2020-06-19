package hindex

func getHIndex(citations []int) int {
	return useBinarySearch(citations)
}

// useBinarySearch time complexity O(log(N)), space complexity O(1)
func useBinarySearch(citations []int) int {
	n := len(citations)
	low, hi := 0, n-1
	for low <= hi {
		mid := low + (hi-low)/2
		if n-mid == citations[mid] {
			return citations[mid]
		} else if n-mid < citations[mid] {
			hi = mid - 1
		} else {
			low = mid + 1
		}
	}
	return n - low
}
