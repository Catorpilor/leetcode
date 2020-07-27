package number

func singleNumber(nums []int) []int {
	return useBit(nums)
}

// useBit time complexity O(N), space complexity O(1)
func useBit(nums []int) []int {
	var diff int
	for i := range nums {
		diff ^= nums[i]
	}
	// pick the rightmost 1 bit
	diff &= -diff
	ans := make([]int, 2)
	for i := range nums {
		if nums[i]&diff == 0 {
			ans[0] ^= nums[i]
		} else {
			ans[1] ^= nums[i]
		}
	}
	return ans
}
