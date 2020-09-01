package pattern

func havePattern(arr []int, m, k int) bool {
	// return useStraight(arr, m, k)
	return useOnepass(arr, m, k)
}

// useStraight time complexity O(N^2), space complexity O(1)
func useStraight(arr []int, m, k int) bool {
	n := len(arr)
	for i := 0; i < n-m+1; i++ {
		j := i + m
		count := 1
		for j < n-m+1 && isEqual(arr[i:i+m], arr[j:j+m]) {
			j += m
			count++
		}
		if count >= k {
			return true
		}
	}
	return false
}

func isEqual(a, b []int) bool {
	n1, n2 := len(a), len(b)
	if n1 != n2 {
		return false
	}
	for i := 0; i < n1; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// useOnepass time complexity O(N), space complexity O(1)
func useOnepass(arr []int, m, k int) bool {
	n := len(arr)
	var cnt int
	for i := 0; i+m < n; i++ {
		if arr[i] != arr[i+m] {
			cnt = 0
		}
		if arr[i] == arr[i+m] {
			cnt++
		}
		if cnt == (k-1)*m {
			return true
		}
	}
	return false
}
