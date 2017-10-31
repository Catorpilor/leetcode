package equal

import "fmt"

func CanPartition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	// 0-1 knapsack problem
	for _, v := range nums {
		for j := target; j >= v; j-- {
			dp[j] = dp[j-v] || dp[j]
		}
	}
	return dp[target]
}

func CanPartition2(nums []int) bool {
	// bits operation
	// bits >> i same as dp[i]
	// for example nums=[2,3,5]
	// when v = 2, bits = 101 represents nums can sum to 0 and 2
	// when v = 3, bits = 101101 represents nums can sum to 0,2,3,5
	// when v = 5, bits = 10110101101 represents nums can sum to 0,2,3,5,7,8,10
	var sum int
	bits := 1
	for _, v := range nums {
		sum += v
		bits |= bits << uint32(v)
	}
	target := uint32(sum / 2)
	return (sum%2 == 0) && (bits>>target)&1 == 1
}
