package targetsum

func maxSubarray(nums []int, target int) int {
	return useHashmap(nums, target)
}

func useHashmap(nums []int, target int) int {
	n := len(nums)
	set := make(map[int]int, n)
	set[0] = -1
	sum := 0
	var ans int
	for i := 0; i < n; i++ {
		sum += nums[i]
		if _, exists := set[sum-target]; exists {
			ans++
			clear(set)
		}
		set[sum] = i
	}
	return ans
}

func clear(set map[int]int) {
	for k := range set {
		delete(set, k)
	}
}
