package tokens

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func maxScore(tokens []int, p int) int {
	// return useTwoPointers(tokens, p)
	return usePrefixSum(tokens, p)
}

// useTwoPointers time complexity O(N*logN), space complexity O(1)
func useTwoPointers(tokens []int, p int) int {
	n := len(tokens)
	sort.Ints(tokens)
	l, r := 0, n-1
	// this is a greedy approach
	var ans, score int
	for l <= r {
		if p >= tokens[l] {
			score++
			p -= tokens[l]
			l++
			ans = utils.Max(ans, score)
		} else if p < tokens[l] && score > 0 {
			p += tokens[r]
			r--
			score--
		} else {
			break
		}
	}
	return ans
}

// usePrefixSum time complexity O(N*logN), space complexity O(N)
func usePrefixSum(nums []int, p int) int {
	n := len(nums)
	sort.Ints(nums)
	ps := make([]int, n)
	copy(ps, nums)
	for i := 1; i < n; i++ {
		ps[i] += ps[i-1]
	}
	if p < ps[0] {
		return 0
	}
	lp := lower_bound(ps, p)
	score, ans := lp, lp
	for i := n - 1; i >= lp && score > 0; i-- {
		p += nums[i]
		// case two
		score--
		pp := lower_bound(ps[:i], p) - lp
		score += pp
		// update left boundary
		lp += pp
		ans = utils.Max(ans, score)
	}
	return ans

}

func binary_search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func lower_bound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
