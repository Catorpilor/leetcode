package longestSubArray

func longestSub(nums []int) int {
	return useTraverse(nums)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func useTraverse(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	var firstOne int
	for i := range nums {
		if nums[i] != 0 {
			break
		}
		firstOne++
	}
	// all 0s
	if firstOne >= n {
		return 0
	}
	curLen, maxLen, prev := 1, 1, 0
	for j := firstOne + 1; j < n; j++ {
		// fmt.Printf("current j:%d\n", j)
		if nums[j] == 1 {
			curLen++
			continue
		}
		// nums[j] == 0 update maxLen
		maxLen = max(maxLen, curLen+prev)
		if j < n-1 && nums[j+1] == 0 {
			// two or more 0s, reset prev
			prev = 0
			// find the next 1
			for zp := j + 1; zp < n; zp++ {
				j = zp
				if nums[zp] != 0 {
					curLen = 1
					// fmt.Printf("update j:%d when facing two or more 0s\n", j)
					break
				}
			}
		} else {
			// maxLen = max(maxLen, prev+curLen)
			prev = curLen
			curLen = 0
		}
	}
	if curLen == n {
		return n - 1
	}
	maxLen = max(maxLen, curLen+prev)
	return maxLen
}
