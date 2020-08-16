package odds

func isStatisfy(nums []int) bool {
	return useOnepass(nums)
}

// useOnepass time complexity O(N), space complexity O(1)
func useOnepass(nums []int) bool {
	var count int
	for i := range nums {
		if nums[i]&1 != 0 {
			count++
			if count == 3 {
				return true

			}
			continue
		}
		count = 0
	}
	return false

}
