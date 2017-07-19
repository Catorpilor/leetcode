package major

func Majority(nums []int) int {
	var major, count int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
			count = 1
		} else {
			if nums[i] == major {
				count += 1
			} else {
				count -= 1
			}
		}
	}
	return major
}
