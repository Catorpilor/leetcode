package x

import (
    "math"

    "github.com/catorpilor/leetcode/utils"
)

func minRes(nums []int, x int) int {
    return useBruteForce(nums, x)
}

func useBruteForce(nums []int, x int) int {
    ans := helper(nums, x, len(nums))
    if ans == math.MaxInt32 {
        return -1
    }
    return ans
}

func helper(nums []int, x, ol int) int {
    n := len(nums)
    if x == 0 {
        return ol - n
    }
    if x < 0 || n == 0 {
        return math.MaxInt32
    }
    return utils.Min(helper(nums[1:], x-nums[0], ol), helper(nums[:n-1], x-nums[n-1], ol))
}
