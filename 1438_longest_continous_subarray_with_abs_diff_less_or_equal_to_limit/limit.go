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

// useDequev2 time complexity O(N), space complexity O(N)
func useDequev2(nums []int, limit int) int {
    // maxd is a monotonic descreasing queue
    // mind is a monotonic increasing queue
    maxd, mind := utils.NewDeque(), utils.NewDeque()
    var i, j int
    n := len(nums)
    var ans int
    for ; j < n; j++ {
        for !maxd.IsEmpty() && nums[j] > maxd.Back().(int) {
            maxd.PopBack()
        }
        for !mind.IsEmpty() && nums[j] < mind.Back().(int) {
            mind.PopBack()
        }
        maxd.PushBack(nums[j])
        mind.PushBack(nums[j])
        for maxd.Front().(int)-mind.Front().(int) > limit {
            // shrink the queue, move the left boundary to the right.
            if maxd.Front().(int) == nums[i] {
                maxd.PopFront()
            }
            if mind.Front().(int) == nums[i] {
                mind.PopFront()
            }
            i++
        }
        ans = utils.Max(ans, j-i+1)
    }
    return ans
}

// useSliceAsDeque time complexity O(N), space complexity O(N)
func useSliceAsDeque(nums []int, limit int) int {
    n := len(nums)
    maxd, mind := make([]int, 0, n), make([]int, 0, n)
    var l, r, ans int
    for ; r < n; r++ {
        l1, l2 := len(maxd), len(mind)
        for l1 > 0 && nums[maxd[l1-1]] < nums[r] {
            maxd = maxd[:l1-1]
            l1--
        }
        for l2 > 0 && nums[mind[l2-1]] > nums[r] {
            mind = mind[:l2-1]
            l2--
        }
        maxd = append(maxd, r)
        mind = append(mind, r)
        if nums[maxd[0]]-nums[mind[0]] > limit {
            l++
            if maxd[0] < l {
                maxd = maxd[1:]
            }
            if mind[0] < l {
                mind = mind[1:]
            }
        }
        ans = utils.Max(ans, r-l+1)
    }
    return ans
}
