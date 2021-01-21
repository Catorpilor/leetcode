package mcs

func mostCompetitive(nums []int, k int) []int {
	return useMonoIncStack(nums, k)
}

// useMonoIncStack time complexity O(N) space complexity O(N)
func useMonoIncStack(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || n == k {
		return nums
	}
	st := make([]int, 0, n)
	ans := make([]int, k)
	for i := range nums {
		nst := len(st)
		for nst > 0 && nums[i] < nums[st[nst-1]] && nst-1+n-i >= k {
			st = st[:nst-1]
			nst--
		}
		if len(st) < k {
			st = append(st, i)
		}
	}
	for i := k - 1; i >= 0; i-- {
		ans[i] = nums[st[i]]
	}
	return ans
}
