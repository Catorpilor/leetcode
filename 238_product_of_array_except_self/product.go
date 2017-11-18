package product

func ProductOfArray(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	preAccu, reversAccu := make([]int, n), make([]int, n)
	preAccu[0], reversAccu[n-1] = nums[0], nums[n-1]
	for i, j := 1, n-2; i < n && j >= 0; i, j = i+1, j-1 {
		preAccu[i], reversAccu[j] = preAccu[i-1]*nums[i], reversAccu[j+1]*nums[j]
	}
	ret := make([]int, n)
	ret[0], ret[n-1] = reversAccu[1], preAccu[n-2]
	for i := 1; i < n-1; i++ {
		ret[i] = preAccu[i-1] * reversAccu[i+1]
	}
	return ret
}

func ProductOfArray2(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	ret := make([]int, n)
	ret[0] = 1
	// ret[1..n-1] stores the accumulated product of nums[0..n-2]
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
