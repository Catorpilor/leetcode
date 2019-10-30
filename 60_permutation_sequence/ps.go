package ps

import (
	"strconv"
	"strings"
)

func getPermutation(n, k int) string {
	if n < 1 {
		return ""
	}
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	res := make([]string, 0, k)
	genRes(&res, nums)
	for j := 1; j < k; j++ {
		nextPermutation(nums)
		genRes(&res, nums)
	}

	// var sb strings.Builder
	// for i := 0; i < n; i++ {
	// 	sb.WriteString(strconv.Itoa(i))
	// }
	// src := sb.String()
	// res := make([]string, 0, k)
	// res = append(res, src)
	// permute(&res, "", 0, src)
	return res[k-1]
}

// based on the top voted discussion
// time complexity O(N) Space complexity O(N)
// the permutations are in order
// for example n=4 [1,2,3,4] k=14
// the permutation sequences are
// 1 + permutation_sequences(2,3,4)  total # is 6 = (n-1)!
// 2 + permutation_sequences(1,3,4)  total # is 6 = (n-1)!
// 3 + permutation_sequences(1,2,4)  total # is 6 = (n-1)!
// 4 + permutation_sequences(1,2,3)  total # is 6 = (n-1)!
// we should know that the 14th sequence should be in the 3rd set
// to get that  use this formular idx = k/(n-1)! = 13/6 = 2, so the first number is 3
// pick 3 out of [1,2,3,4] so the array becomes [1,2,4]
// then we update k, k = k - idx*(n-1)! # of skipped sequences
// then we continue to find the next bit.

func getKthPermutation(n, k int) string {
	if n < 1 {
		return ""
	}
	factorial := make([]int, n+1)
	factorial[0] = 1
	factor := 1
	for i := 1; i < n+1; i++ {
		factor *= i
		factorial[i] = factor
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	k--
	var st strings.Builder
	for i := 1; i <= n; i++ {
		idx := k / factorial[n-i]
		st.WriteString(strconv.Itoa(nums[idx]))
		nums = append(nums[:idx], nums[idx+1:]...)
		k -= idx * factorial[n-i]
	}
	return st.String()
}

// just call nextPermutation k times
// src are monotonically increasing "123456...n"

func nextPermutation(nums []int) {
	i, j := len(nums)-2, len(nums)-1

	for i >= 0 && nums[i+1] < nums[i] {
		i--
	}
	if i >= 0 {
		for j >= 0 && nums[j] < nums[i] {
			j--
		}
	}
	// swap nums[i], nums[j]
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums, i+1, len(nums)-1)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// func nextPermutation(src string) string{
// 	n := len(src)
// 	i := n - 1
// 	for i >=0 && src[i-1] > src[i] {
// 		i--
// 	}
// 	if i > = 0 {
// 		j :=
// 	}
// }

func genRes(res *[]string, nums []int) {
	var st strings.Builder
	for i := range nums {
		st.WriteString(strconv.Itoa(nums[i]))
	}
	*res = append(*res, st.String())
}

// func permute(res *[]string, prefix string, pos int, s string) {
// 	if pos == len(s) {
// 		*res = append(*res, prefix)
// 		return
// 	}
// 	for i := pos; i < len(s); i++ {

// 	}
// }
