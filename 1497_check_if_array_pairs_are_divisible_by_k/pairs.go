package pairs

func canArrange(arr []int, k int) bool {
	return useModule(arr, k)
}

//useModule time complexity O(N), space complexity O(N)
func useModule(arr []int, k int) bool {
	n := len(arr)
	if n%2 != 0 {
		return false
	}
	if k == 1 {
		return true
	}
	//remains store the occurreneces of the remainder of arr[i]
	remains := make([]int, k)
	for i := range arr {
		rem := arr[i] % k
		if rem < 0 {
			rem += k
		}
		remains[rem]++
	}
	for i := 1; i < k; i++ {
		if remains[i] != remains[k-i] {
			// can not be paired
			return false
		}
	}
	return remains[0]%2 == 0
}
