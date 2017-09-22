package ms

func MissingNumber(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res ^= i
		res ^= nums[i]
	}
	return res
}
