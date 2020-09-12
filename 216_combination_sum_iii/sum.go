package sum

func combinationSum(k, n int) [][]int {
	// backtracking
	return useBackTracking(k, n)

}

func useBackTracking(k, n int) [][]int {
	var ans [][]int
	comb := make([]int, 0, k)
	helper(&ans, comb, 1, n, k)
	return ans
}

func helper(res *[][]int, temp []int, start, n, k int) {
	if k < 0 {
		return
	}
	if k == 0 {
		if n == 0 {
			bak := make([]int, len(temp))
			copy(bak, temp)
			*res = append(*res, bak)
		}
	} else {
		for i := start; i <= 9 && n >= i; i++ {
			lt := len(temp)
			temp = append(temp, i)
			helper(res, temp, i+1, n-i, k-1)
			temp = temp[:lt]
		}
	}
}
