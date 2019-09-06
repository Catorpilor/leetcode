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
    // brute force
    res := make([]int, n-k+1)
    for i := 0; i < n-k+1; i++ {
        curMax := math.MinInt32
        for j := i; j < i+k; j++ {
            curMax = utils.Max(curMax, nums[j])
        }
        res[i] = curMax
    }

    // pMax := make([]int, len(nums))
    // pMax[0] = nums[0]
    // curMax := nums[0]
    // for i := 1; i < n; i++ {
    //     if nums[i] > curMax {
    //         curMax = nums[i]
    //     }
    //     pMax[i] = curMax
    // }
    // fmt.Printf("pMax is %v\n", pMax)
    return res

}
