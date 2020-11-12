package nlis

func findNumberOfLIS(nums []int) int {
	return useDP(nums)
}

func useDPw(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	// dp[i] represents # of lis ended with nums[i]
	// mLen[i] represents the length of LIS ended with nums[i]
	dp, mLen := make([]int, n), make([]int, n)
	for i := range mLen {
		dp[i], mLen[i] = 1, 1
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if mLen[i] < mLen[j]+1 {
					mLen[i] = mLen[j] + 1
					dp[i] = dp[j]
				} else if mLen[i] == mLen[j]+1 {
					dp[i] += dp[j]
				}
			}
		}
	}
	var ret, curMax int
	for i := range mLen {
		if curMax < mLen[i] {
			ret = dp[i]
			curMax = mLen[i]
		} else if curMax == mLen[i] {
			ret += dp[i]
		}
	}
	return ret
}
