package choc

func divideChocolate(chocolate []int, k int) int {
	return useBinarySearch(chocolate, k)
}

// useBinarySearch time complexity O(N*log(M)) where M = max(arr) - min(arr), space complexity O(1)
func useBinarySearch(chocolate []int, k int) int {
	n := len(chocolate)
	if n < 1 {
		return 0
	}
	minVal, preSum := chocolate[0], chocolate[0]
	for i := 1; i < n; i++ {
		preSum += chocolate[i]
		if chocolate[i] < minVal {
			minVal = chocolate[i]
		}
	}
	l, r := minVal, preSum
	for l < r {
		mid := r - (r-l)/2
		if count(chocolate, mid) < k {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

func count(arr []int, target int) int {
	sum, ans := 0, 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum >= target {
			ans++
			sum = 0
		}
	}
	return ans
}
