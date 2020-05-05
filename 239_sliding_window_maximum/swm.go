package swm

import (
	"fmt"
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	// return bf(nums, k, n)
	return dp(nums, k, n)
}

func bf(nums []int, k, n int) []int {
	res := make([]int, n-k+1)
	for i := 0; i < n-k+1; i++ {
		curMax := math.MinInt32
		for j := i; j < i+k; j++ {
			curMax = utils.Max(curMax, nums[j])
		}
		res[i] = curMax
	}
	return res
}

func dp(nums []int, k, n int) []int {
	// divide nums into blocks of k elements
	// the last block could contain less than k elements
	// left[i] from block_start to i the max element
	// right[j] from block_end to j the max element
	//      [1, 3 -1, -3, 5, 3, 6, 7] k = 3
	// left [1, 3, 3, -3, 5, 5, 6, 7]
	// right[3, 3,-1, 5,  5, 3, 7, 7]
	// so for slide window i, j
	// maxVal = max(left[j], right[i])
	left, right := make([]int, n), make([]int, n)
	left[0], right[n-1] = nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		if i%k == 0 {
			left[i] = nums[i]
		} else {
			left[i] = utils.Max(left[i-1], nums[i])
		}

		j := n - i - 1
		if j%k == 0 {
			right[j] = nums[j]
		} else {
			right[j] = utils.Max(right[j+1], nums[j])
		}
	}
	res := make([]int, n-k+1)
	for i := 0; i < n-k+1; i++ {
		res[i] = utils.Max(left[i+k-1], right[i])
	}
	return res
}

// useDeque time complexity O(N), space complexity O(N)
func useDeque(nums []int, k int) []int {
	n := len(nums)
	if n < 1 || k <= 0 {
		return nil
	}
	res := make([]int, n-k+1)
	ri := 0
	dq := utils.NewDeque()
	for i := 0; i < n; i++ {
		// fmt.Printf("idx:%d, deq: %s\n", i, dq.String())
		// remove numbers out of range k
		// in one iteration we only pop one element and push back 1.
		if !dq.IsEmpty() && dq.Front().(int) < i-k+1 {
			// fmt.Printf("before remove from front, got element %d\n", dq.Front().(int))
			dq.PopFront()
		}
		// remove smaller numbers in k range as they are useless
		// Now only those elements within [i-(k-1),i] are in the deque.
		// We then discard elements smaller than a[i] from the tail.
		// This is because if a[x] <a[i] and x<i, then a[x] has no chance to be the "max" in [i-(k-1),i],
		// or any other subsequent window: a[i] would always be a better candidate.
		for !dq.IsEmpty() && nums[dq.Back().(int)] < nums[i] {
			// fmt.Printf("before remove from back, got element %d\n", dq.Back().(int))
			dq.PopBack()
		}
		dq.PushBack(i)
		fmt.Printf("after push back i:%d, deq: %s\n", i, dq.String())
		if i >= k-1 {
			res[ri] = nums[dq.Front().(int)]
			ri++
		}
	}
	return res
}
