package gen

import "github.com/catorpilor/leetcode/utils"

func getMaximumGen(n int) int {
	return useIter(n)
}

// useIter time complexity O(N), space complexity O(N)
func useIter(n int) int {
	if n <= 1 {
		return n
	}
	res := make([]int, n+1)
	res[1] = 1
	ans := 0
	for i := 1; i <= n; i++ {
		if i<<1 <= n {
			res[i<<1] = res[i]
		}
		if i<<1+1 <= n {
			res[i<<1+1] = res[i] + res[i+1]
		}
		ans = utils.Max(ans, res[i])
	}
	return ans
}
