package pps

// prevPerm time complexity O(N), space complexity O(1)
func prevPerm(nums []int) []int {
    n := len(nums)
    if n <= 1 {
        return nums
    }
    idx := -1
    // find the largest i that nums[i] > nums[i+1]
    for i := n - 1; i >= 1; i-- {
        if nums[i-1] > nums[i] {
            idx = i - 1
            break
        }
    }
    if idx == -1 {
        // not found
        return nums
    }
    // traverse from right, find the first i that satisifies nums[i] < nums[idx] && nums[i] != nums[i-1]
    for i := n - 1; i > idx; i-- {
        if nums[i] < nums[idx] && nums[i] != nums[i-1] {
            // swap
            nums[i], nums[idx] = nums[idx], nums[i]
            break
        }
    }
    return nums
}
