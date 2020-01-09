package splk

func numsOfSubarray(nums []int, k int) int {
	return mine(nums, k)
}

func twoPointers(nums []int, k int) int {
	n := len(nums)
	p := 1
	l, r := 0, 0
	var res int
	for r < n {
		p *= nums[r]
		for l <= r && p >= k {
			p /= nums[l]
			l++
		}
		res += r - l + 1
		r++
	}
	return res
}

// mine inspired by twoPointers, did some modification
// time complexity O(n^2), space complexity is O(n)
func mine(nums []int, k int) int {
	n := len(nums)
	var res int
	if n < 1 {
		return res
	}
	// sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	opts := make([]int, n)
	for i := range opts {
		opts[i] = i
	}
	// fmt.Printf("before parse: opt: %v\n", opts)
	prodSofar, l := 1, 0
	for i := 0; i < n; i++ {
		prodSofar *= nums[i]
		if prodSofar < k {
			if i > 0 && opts[i-1] != -1 {
				opts[i] = opts[i-1]
			}
			continue
		}
		for l <= i && prodSofar >= k {
			prodSofar /= nums[l]
			l++
		}
		// fmt.Printf("current i:%d, temp: %d, k:%d \n", i, temp, k)
		if l > i {
			opts[i] = -1
		} else {
			opts[i] = l
		}

	}
	// fmt.Printf("opts: %v\n", opts)
	for i := range opts {
		if opts[i] == -1 {
			continue
		}
		res += i - opts[i] + 1
	}
	return res
}
