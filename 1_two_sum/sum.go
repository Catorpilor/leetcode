package sum

func twoSums(nums []int, target int) []int {
	return useHashmap(nums, target)
}

// useHashmap time complexity O(N), space complexity O(N)
func useHashmap(nums []int, target int) []int {
	ans := make([]int, 0, 2)
	n := len(nums)
	if n < 2 {
		return ans
	}
	// set's key is  target-nums[i] and value is the position of nums[i]
	set := make(map[int]int, n)
	for i, num := range nums {
		if pos, exists := set[num]; exists {
			ans = append(ans, pos, i)
			return ans
		} else {
			set[target-num] = i
		}
	}
	return ans
}
