package shuffle

func shuf(nums []int, n int) []int {
	return useSlice(nums, n)
}

func useSlice(nums []int, n int) []int {
	res := make([]int, 0, n)
	prev, last := nums[:n], nums[n:]
	for i := 0; i < n; i++ {
		res = append(res, prev[i], last[i])
	}
	return res
}
