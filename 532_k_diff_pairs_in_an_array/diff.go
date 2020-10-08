package diff

func findPairs(nums []int, k int) int {
	return useHashmap(nums, k)
}

// useHashmap time complexity O(N), space complexity O(N)
func useHashmap(nums []int, k int) int {
	n := len(nums)
	set := make(map[int]int, n)
	for _, num := range nums {
		set[num]++
	}
	var ans int
	for key := range set {
		if (k > 0 && set[key+k] > 0) || (k == 0 && set[key] > 1) {
			ans++
		}
	}
	return ans
}
