package shuffle

func shuf(nums []int, n int) []int {
	return useSlice(nums, n)
}

// useSlice time complexity O(N), space complexity O(1)
func useSlice(nums []int, n int) []int {
	res := make([]int, 0, n)
	prev, last := nums[:n], nums[n:]
	for i := 0; i < n; i++ {
		res = append(res, prev[i], last[i])
	}
	return res
}

func useMath(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i := range nums {
		if i%2 == 0 {
			res[i] = nums[i/2]
		} else {
			res[i] = nums[n+i/2]
		}
	}
	return res
}
