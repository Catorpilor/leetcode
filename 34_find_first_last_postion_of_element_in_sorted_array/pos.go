package pos

func searchRange(arr []int, target int) []int {
	l := lowerBound(arr, target)
	r := lowerBound(arr, target+1)
	if l < r {
		return []int{l, r - 1}
	}
	return []int{-1, -1}
}

// lowberBound time complexity O(lgN), space complexity O(1)
func lowerBound(arr []int, target int) int {
	n := len(arr)
	low, hi := 0, n-1
	for low <= hi {
		mid := low + (hi-low)>>1
		if arr[mid] >= target {
			hi = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
