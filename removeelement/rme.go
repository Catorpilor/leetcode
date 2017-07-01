package removeelement

func RemoveElement(nums []int, val int) int {
	l := len(nums)
	ret := 0
	i := 0
	for i < l {
		if nums[i] == val {
			// exchange with the last element
			temp := nums[l-1]
			nums[l-1] = nums[i]
			nums[i] = temp
			nums = nums[:l-1]
			l = len(nums)
		} else {
			ret++
			i++
		}
	}
	return ret
}
