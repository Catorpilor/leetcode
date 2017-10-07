package desc

func Desc(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	var cnt int
	for i := 1; i< len(nums) && cnt <= 1;i++ {
		if nums[i-1] > nums[i] {
			cnt += 1
			if i - 2 < 0 || nums[i-2] <= nums[i] {
				nums[i-1] = nums[i]
			}else {
				nums[i] = nums[i-1]
			}
		}
	}
	return cnt <= 1
	// curMax := nums[0]
	// for i := 0; i<len(nums)-1; i++ {
	// 	if nums[i] > nums[i+1] {
	// 		count += 1
	// 		if count > 1 {
	// 			return false
	// 		}
	// 	}else {
	// 		curMax = max(curMax,nums[i+1])
	// 	}
	// }
	// return true
}

func max(a,b int) int{
	if a < b {
		return b
	}
	return a
}