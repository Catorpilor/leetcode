package swm

import (
    "math"

    "github.com/catorpilor/leetcode/utils"
)

func MaxSlidingWindow(nums []int, k int) []int {
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
        if (j+1)%k == 0 {
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
