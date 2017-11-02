package combine

func Combination(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 1; j <= target; j++ {
		for i := range nums {
			if nums[i] <= j {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
