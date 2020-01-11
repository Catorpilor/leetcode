package css

func checkSubArraySum(nums []int, k int) bool {
	return bruteForce(nums, k)
}

func bruteForce(nums []int, k int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	var res bool
	var sum int
	for i := range nums {
		sum = 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum == k || (k != 0 && sum%k == 0) {
				if j > i {
					// at least 2 that sum up to k
					return true
				}
			}
		}
	}
	return false
}
