package flood

func avoidFlood(rain []int) []int {
	// return straight(rain)
	return useLowerBound(rain)
}

// useLowerBound time complexity O(N*log(N)), space complexity O(N)
func useLowerBound(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	dryDays := make([]int, 0, n)
	fullLakes := make(map[int]int, n)
	for i, v := range rains {
		if v == 0 {
			dryDays = append(dryDays, i)
			ans[i] = 1
		} else {
			if d, exists := fullLakes[v]; exists {
				lb := lowerBound(dryDays, d)
				if lb != -1 {
					ans[dryDays[lb]] = v
					dryDays = append(dryDays[:lb], dryDays[lb+1:]...)
				} else {
					return []int{}
				}
			}
			fullLakes[v] = i
			ans[i] = -1
		}
	}
	return ans
}

// lowerBount time complexity O(log(N))
func lowerBound(arr []int, target int) int {
	n := len(arr)
	l, r := 0, n-1
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] > target {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}
