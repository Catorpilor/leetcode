package nlis

func findNumberOfLIS(nums []int) int {
	return useDP(nums)
}

func useDP(nums []int) int {
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

// useSimpleDP time complexity O(N^2), space complexity O(N)
func useSimpleDP(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	// cnt[i] stores the number of lis ends with nums[i]
	// lens[i] stores the length of lis ends with nums[i]
	cnt, lens := make([]int, n), make([]int, n)
	var ans, maxLen int
	for i := 0; i < n; i++ {
		cnt[i], lens[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// a possible lis
				if lens[i] == lens[j]+1 {
					cnt[i] += cnt[j]
				}
				if lens[i] < lens[j]+1 {
					lens[i] = lens[j] + 1
					cnt[i] = cnt[j]
				}
			}
		}
		if maxLen == lens[i] {
			ans += cnt[i]
		}
		if maxLen < lens[i] {
			ans = cnt[i]
			maxLen = lens[i]
		}
	}
	return ans
}
