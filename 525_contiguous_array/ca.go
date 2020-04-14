package ca

import "github.com/catorpilor/leetcode/utils"

func maxContiguousArray(nums []int) int {
    return useBruteForce(nums)
}

// useBruteForce time complexity O(N^2), space complexity O(1)
func useBruteForce(nums []int) int {
    n := len(nums)
    var ans int
    count := make([]int, 2)
    for i := 0; i < n-1; i++ {
        count[nums[i]] = 1
        for j := i + 1; j < n; j++ {
            count[nums[j]]++
            if count[0] == count[1] {
                ans = utils.Max(ans, j-i+1)
            }
        }
    }
    return ans
}
