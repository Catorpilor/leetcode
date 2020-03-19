package product

func productExceptSelf(nums []int) []int {
	return useAddSpace(nums)
}

// useAddSpace time complexity O(N), space complexity O(N)
func useAddSpace(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	preAccu, reversAccu := make([]int, n), make([]int, n)
	copy(preAccu, nums)
	copy(reversAccu, nums)
	for i := 1; i < n; i++ {
		preAccu[i] *= preAccu[i-1]
		reversAccu[n-i-1] *= reversAccu[n-i]
	}
	ans := make([]int, n)
	ans[0] = reversAccu[1]
	ans[n-1] = preAccu[n-2]
	for i := 1; i < n-1; i++ {
		ans[i] = preAccu[i-1] * reversAccu[i+1]
	}
	return ans
}

// useConstantSpace time complexity O(N), space complexity O(1)
func useConstantSpace(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	ret := make([]int, n)
	ret[0] = 1
	// initialize ret, ret[i] = the product of nums[0]..nums[i-1]
	for i := 1; i < n; i++ {
		ret[i] = ret[i-1] * nums[i-1]
	}
	right := 1
	for j := n - 1; j >= 0; j-- {
		ret[j] *= right
		right *= nums[j]
	}
	return ret
}
