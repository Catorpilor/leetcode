package max

func equal(a int, b *int) bool {
	if b == nil {
		return false
	}
	return a == *b
}

func ThirdMax(nums []int) int {
	var max1, max2, max3 *int
	for i := 0; i < len(nums); i++ {
		if equal(nums[i], max1) || equal(nums[i], max2) || equal(nums[i], max3) {
			continue
		}
		if max1 == nil || nums[i] > *max1 {
			max3 = max2
			max2 = max1
			max1 = &nums[i]
		} else if max2 == nil || nums[i] > *max2 {
			max3 = max2
			max2 = &nums[i]
		} else if max3 == nil || nums[i] > *max3 {
			max3 = &nums[i]
		}
	}
	if max3 == nil {
		return *max1
	}
	return *max3
}
