package as

func NumberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	dp := make([][]bool, n)
	prev, ret := nums[1]-nums[0], 0
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n-2; i++ {
		if nums[i+2]-nums[i+1] == prev {
			dp[i][i+2] = true
			ret += 1
		}
		prev = nums[i+2] - nums[i+1]
	}
	// for length >= 4
	for l := 4; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			dp[i][j] = dp[i][j-1] && dp[i+1][j]
			if dp[i][j] {
				ret += 1
			}
		}
	}
	return ret
}

func NumberOfArithmeticSlices2(nums []int) int {
	var cur, sum int
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			cur += 1
			sum += cur
		} else {
			cur = 0
		}
	}
	return sum
}

func NumberOfArithmeticSlices3(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	dp := make([]int, n)
	// dp[i] represents the number of arithmetic slices ending at index i
	if nums[2]-nums[1] == nums[1]-nums[0] {
		dp[2] = 1
	}
	ret := dp[2]
	for i := 3; i < n; i++ {
		// if A[i-2], A[i-1], A[i] are arithmetic, then the number of arithmetic slices ending with A[i] (dp[i])
		// equals to:
		//      the number of arithmetic slices ending with A[i-1] (dp[i-1], all these arithmetic slices appending A[i] are also arithmetic)
		//      +
		//      A[i-2], A[i-1], A[i] (a brand new arithmetic slice)
		// it is how dp[i] = dp[i-1] + 1 comes
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
		ret += dp[i]
	}
	return ret
}
