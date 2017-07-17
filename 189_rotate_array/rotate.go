package rotate

func Rotate(nums []int, k int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	k = k % l
	rotate(nums[:l-k])
	rotate(nums[l-k:])
	rotate(nums)
	return nums
}

func rotate(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
