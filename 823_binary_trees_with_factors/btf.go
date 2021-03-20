package btf

import (
	"sort"
)

func numberOfFactorBinaryTree(arr []int) int {
	return useDP(arr)
}

// useDP time complexity O(N^2) space complexity O(N)
func useDP(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return n
	}
	sort.Ints(arr)
	dp := make([]int, n)
	hset := make(map[int]int, n)
	for i := range dp {
		dp[i] = 1
		hset[arr[i]] = i
	}
	mod := 1_000_000_007
	for i := 0; i<n; i++ {
		for j:=0; j<i; j++ {
			if arr[i] % arr[j] == 0 {
				// arr[j] is the left child, arr[i] is the root
				right := arr[i] / arr[j]
				if v, exists := hset[right]; exists {
					dp[i] = (dp[i] + dp[j]*dp[v]) % mod
				}
			}
		}
	}
	var ans int64
	for i := range dp {
		ans += int64(dp[i])
	}
	return int(ans % int64(mod))
}

