package maxsub

import (
    "math"

    "github.com/catorpilor/leetcode/utils"
)

func maxSubSum(nums []int) int {
    return twoPasses(nums)
}

// twoPasses time complexity O(N), space complexity O(1)
func twoPasses(nums []int) int {
    n := len(nums)
    first := useKadane(nums, n)
    preSum := 0
    for i := range nums {
        preSum += nums[i]
        nums[i] = -nums[i]
    }
    preSum += useKadane(nums, n)
    if first > 0 {
        return utils.Max(first, preSum)
    }
    return first
}

// useKadane use kadanes's algorightms to find the max sum subarray
// but dit some modification, since we want non-empty subarray.
func useKadane(nums []int, n int) int {
    var maxEndingHere int
    maxSofar := math.MinInt32
    for i := 0; i < n; i++ {
        maxEndingHere += nums[i]
        if maxEndingHere > maxSofar {
            maxSofar = maxEndingHere
        }
        if maxEndingHere < 0 {
            maxEndingHere = 0
        }
    }
    return maxSofar
}
