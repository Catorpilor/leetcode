package ap

import "sort"

func arrayPairSum(nums []int) int {
    return sorted(nums)
}

// sorted sort nums first and sum the even index up
// time complexity is O(NlgN), space complexity is O(1)
func sorted(nums []int) int {
    n := len(nums)
    if n < 2 || n%2 != 0 {
        return 0
    }
    // sort nums
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    var ret int
    for i := 0; i < n; i += 2 {
        ret += nums[i]
    }
    return ret
}
