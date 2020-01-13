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

// countSort needs to konw the range of element
// time complexity is O(N), Space complexity is o(N)
func countSort(nums []int) int {
    n := len(nums)
    if n < 2 || n%2 != 0 {
        return 0
    }
    // nums[i] in [-10000,10000] range
    exists := make([]int, 20001)
    for i := range nums {
        // make a count
        exists[10000+nums[i]]++
    }
    flag := false
    var res int
    // iterate exists which just like nums is sorted
    // for example nums = [2,1,3,4,2,4]
    // exists should be exists[:10000] = 0,
    // exists[10001] = 1
    // exists[10002] = 2
    // exists[10003] = 1
    // exists[10004] = 2
    // sorted nums [1,2,2,3,4,4]
    // just like the sorted solution, pair[1,2] we pick 1, pair[2,3] pick 2
    for j := range exists {
        for exists[j] > 0 {
            // use flag to pick the even indexed [0,2,4,....]
            if !flag {
                res += j - 10000
            }
            flag = !flag
            exists[j]--
        }
    }
    return res
}
