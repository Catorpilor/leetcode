package sum

import "sort"

func CombinationSum(nums []int, target int) [][]int {
	// dp is too compicated to get the detailed sub sets
	// // dp[i][j] represents for target j how many ways to sum.
	// n := len(nums)
	// if n < 1 {
	// 	return nil
	// }
	// dp := make([][]int, n)
	// for i := range dp {
	// 	dp[i] = make([]int, target+1)
	// 	dp[i][0] = 1
	// }
	// for i := 1; i <= target; i++ {
	// 	if i%nums[0] == 0 {
	// 		dp[0][i] = 1
	// 	}
	// }
	// for i := 1; i < n; i++ {
	// 	for j := 1; j <= target; j++ {
	// 		if nums[i] > j {
	// 			dp[i][j] = dp[i-1][j]
	// 		} else {
	// 			dp[i][j] = dp[i-1][j] + dp[i][j-nums[i]]
	// 		}
	// 	}
	// }
	// fmt.Println(dp)

	// return nil
	// here we use backtrack
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ret := make([][]int, 0)
	temp := make([]int, 0)
	backtrace(&ret, temp, nums, target, 0)
	return ret
}

func backtrace(res *[][]int, temp, nums []int, target, start int) {
	if target < 0 {
		return
	} else if target == 0 {
		tt := make([]int, len(temp))
		copy(tt, temp)
		*res = append(*res, tt)
	} else {
		for i := start; i < len(nums) && target >= nums[i]; i++ {
			l := len(temp)
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			backtrace(res, temp, nums, target-nums[i], i+1)
			temp = temp[:l] // maybe not efficient
		}
	}
}
