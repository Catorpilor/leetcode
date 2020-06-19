package bouquet

func minDays(bloomDay []int, m, k int) int {
	return useBinarySearch(bloomDay, m, k)
}

// useBinarySearch time complexity O(N*log(K)), space complexity O(1)
func useBinarySearch(bloomDay []int, m, k int) int {
	n := len(bloomDay)
	if m*k > n {
		return -1
	}
	low, hi := 1, int(1e9)
	for low < hi {
		mid := low + (hi-low)/2
		var bloomedFlower, bouq int
		for j := 0; j < n; j++ {
			if bloomDay[j] > mid {
				bloomedFlower = 0
			} else {
				bloomedFlower++
				if bloomedFlower >= k {
					bouq++
					bloomedFlower = 0
				}
			}
		}
		if bouq < m {
			low = mid + 1
		} else {
			hi = mid
		}
	}
	return low
}
