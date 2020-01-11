package css

func checkSubArraySum(nums []int, k int) bool {
	return bruteForce(nums, k)
}

// bruteForce time complexity O(n^2) space complexity O(1)
func bruteForce(nums []int, k int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
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

// preSumHash time complexity o(N), space complexity O(N)
func preSumHash(nums []int, k int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	// key is the sum%k, value is index
	// based on this math equation
	// sum1 % k = x ----- eq1
	// sum2 % k = x ----- eq2
	// then (sum1-sum2) % k = x - x = 0 ------ eq3
	hash := make(map[int]int, n)
	// edge case
	hash[0] = -1
	var sum int
	for i := range nums {
		sum += nums[i]
		if k != 0 {
			sum %= k
		}
		if v, exists := hash[sum]; exists {
			// found one based on eq3
			if i-v > 1 {
				return true
			}
		} else {
			hash[sum] = i
		}
	}
	return false
}
