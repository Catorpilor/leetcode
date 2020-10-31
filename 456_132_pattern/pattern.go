package pattern

func findPattern(nums []int) bool {
	return useBruteForce(nums)
}

// useBruteForce time complexity O(N^3), space complexity O(1)
func useBruteForce(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i] < nums[k] && nums[j] > nums[k] {
					return true
				}
			}
		}
	}
	return false
}
