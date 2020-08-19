package candies

func distributeCandies(total, num int) []int {
	return useMath(total, num)
}

// useMath time complexity O(1), space complexity O(N)
func useMath(total, num int) []int {
	// it's a arithmetic sequence
	// we find the max `n` which statisfy the n*(n+1) < 2*total
	ans := make([]int, num)
	upBound := 1
	cur, pre := 1, 1
	for cur < total {
		upBound++
		pre = cur
		cur = upBound * (upBound + 1) / 2
	}
	for i := 1; i < upBound; i++ {
		pos := (i - 1) % num
		ans[pos] += i
	}
	ans[(upBound-1)%num] += total - pre
	return ans
}
