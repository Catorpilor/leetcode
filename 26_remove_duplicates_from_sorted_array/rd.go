package rd

func removeDups(nums []int) int {
	return useTraverse(nums)
}

// useTraverse time complexity O(N), space complexity O(1)
func useTraverse(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	prev, ans := nums[0], 1
	for i := 1; i < n; i++ {
		if nums[i] != prev {
			prev = nums[i]
			nums[ans], nums[i] = nums[i], nums[ans]
			ans++
		}
	}
	return ans
}
