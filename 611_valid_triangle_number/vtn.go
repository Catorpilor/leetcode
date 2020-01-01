package vtn

import "sort"

func triangleNumbers(nums []int) int {
    return bruteForce(nums)
}

func bruteForce(nums []int) int {
    var res int
    n := len(nums)
    if n < 3 {
        return res
    }
    // sort nums O(nlgn)
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] <= nums[j]
    })
    // time complexity O(n^3)
    for i := 0; i < n-2; i++ {
        for j := i + 1; j < n-1; j++ {
            for k := j + 1; k < n; k++ {
                if nums[i]+nums[j] > nums[k] {
                    res++
                } else {
                    break
                }
            }
        }
    }

    return res
}

// func helper(nums []int, res *int, pos int) {

// }
