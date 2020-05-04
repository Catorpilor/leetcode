package limit

func longestSubarray(nums []int, limit int) int {
    return usebf(nums, limit)
}

// usebf time complexity O(N^3), space complexity O(1)
func usebf(nums []int, limit int) int {
    ans := 1
    n := len(nums)
    for k := 2; k <= n; k++ {
        for i := 0; i+k <= n; i++ {
            // subarray nums[i:i+k]
            lmin, lmax := nums[i], nums[i]
            for j := i + 1; j < i+k; j++ {
                if nums[j] > lmax {
                    lmax = nums[j]
                }
                if nums[j] < lmin {
                    lmin = nums[j]
                }
            }
            if lmax-lmin <= limit {
                ans = max(ans, k)
            }
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
