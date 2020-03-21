package sum

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func digitSum(nums []int) int {
	// return useBruteForce(nums)
	return useOneTraversal(nums)
}

// useBruteForce time complexity O(K*N), space complexity O(N)
func useBruteForce(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return -1
	}
	// hset stores all the nums with same digit sum
	hset := make(map[int][]int, n)
	for _, num := range nums {
		ds := calculateSum(num)
		hset[ds] = append(hset[ds], num)
	}
	ans := -1
	for _, v := range hset {
		nv := len(v)
		if nv < 2 {
			continue
		} else if nv == 2 {
			ans = utils.Max(ans, v[0]+v[1])
		} else {
			sort.Ints(v)
			ans = utils.Max(ans, v[nv-2]+v[nv-1])
		}
	}
	return ans

}

// calculateSum time complexity O(K), space complexity O(1)
func calculateSum(num int) int {
	var ans int
	num = utils.Abs(num)
	for num != 0 {
		ans += num % 10
		num /= 10
	}
	return ans
}

// useOneTraversal time complexity O(K*N), space complexity O(N)
func useOneTraversal(nums []int) int {
	n := len(nums)
	if n < 2 {
		return -1
	}
	ans := -1
	// set just stores the max num of the same digit sum we've seen so far
	set := make(map[int]int, n)
	for _, num := range nums {
		ds := calculateSum(num)
		if v, exists := set[ds]; exists {
			ans = utils.Max(ans, num+v)
			set[ds] = utils.Max(v, num)
		} else {
			set[ds] = num
		}
	}
	return ans
}
