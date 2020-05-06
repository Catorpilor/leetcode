package limit

import "github.com/catorpilor/leetcode/utils"

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

// useDeque time complexity O(N), space complexity O(N)
func useDeque(nums []int, limit int) int {
    maxd, mind := utils.NewDeque(), utils.NewDeque()
    var i, j int
    n := len(nums)
    for ; j < n; j++ {
        for !maxd.IsEmpty() && nums[j] > maxd.Back().(int) {
            maxd.PopBack()
        }
        for !mind.IsEmpty() && nums[j] < mind.Back().(int) {
            mind.PopBack()
        }
        maxd.PushBack(nums[j])
        mind.PushBack(nums[j])
        if maxd.Front().(int)-mind.Front().(int) > limit {
            if maxd.Front().(int) == nums[i] {
                maxd.PopFront()
            }
            if mind.Front().(int) == nums[i] {
                mind.PopFront()
            }
            i++
        }
    }
    return j - i
}
