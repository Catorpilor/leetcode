package jump

import "github.com/catorpilor/leetcode/utils"

func canJump(arr []int) bool {
    return useGreedyForward(arr)
}

// useGreedyForward time complexity O(N), space complexity O(1)
func useGreedyForward(nums []int) bool {
    n := len(nums)
    if n < 1 {
        return false
    }
    // the maxPos can get so far
    maxPos := 0
    for i := 0; i <= maxPos; i++ {
        // update the right boundary if necessary.
        maxPos = utils.Max(maxPos, i+nums[i])
        if maxPos >= n-1 {
            return true
        }
    }
    return false
}
