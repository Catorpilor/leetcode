package major

func Majority(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	// O(n) space O(n) time
	hash := make(map[int]int)
	for i := range nums {
		hash[nums[i]] += 1
	}
	ret := make([]int, 0, n)
	for k, v := range hash {
		if v > n/3 {
			ret = append(ret, k)
		}
	}
	return ret
}

func Majority2(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	// boyer-moore voting algorithm
	// want majority elements more than [n/3]
	// so there are <=2 elements
	candi1, candi2, count1, count2 := 0, 1, 0, 0
	for i := range nums {
		if nums[i] == candi1 {
			count1 += 1
		} else if nums[i] == candi2 {
			count2 += 1
		} else if count1 == 0 {
			candi1, count1 = nums[i], 1
		} else if count2 == 0 {
			candi2, count2 = nums[i], 1
		} else {
			count1, count2 = count1-1, count2-1
		}
	}
	count1, count2 = 0, 0
	for i := range nums {
		if nums[i] == candi1 {
			count1 += 1
		}
		if nums[i] == candi2 {
			count2 += 1
		}
	}
	ret := make([]int, 0, 2)
	if count1 > n/3 {
		ret = append(ret, candi1)
	}
	if count2 > n/3 {
		ret = append(ret, candi2)
	}
	return ret

}
