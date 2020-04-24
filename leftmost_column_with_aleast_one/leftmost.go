package leftmost

import "github.com/catorpilor/leetcode/utils"

// board is at most 100*100
// restricions: calling board[i][j] no more than 1000 times.

func leftMostColumnWithOne(board [][]int) int {
    // return useBinarySearch(board)
    return useTraverse(board)
}

// useBinarySearch time complexity O(M*lgN), space complexity O(1)
func useBinarySearch(board [][]int) int {
    // board's row element is non-decresing order.
    m := len(board)
    ans := -1
    if m < 1 {
        return ans
    }
    n := len(board[0])
    if n < 1 {
        return ans
    }
    for i := 0; i < m; i++ {
        tmp := lowerBoundOfOne(board[i])
        if tmp != -1 {
            if ans != -1 {
                ans = utils.Min(ans, tmp)
            } else {
                ans = tmp
            }
        }
    }
    return ans
}

func lowerBoundOfOne(nums []int) int {
    ans := -1
    l, r := 0, len(nums)-1
    for l <= r {
        mid := l + (r-l)/2
        if nums[mid] != 1 {
            l = mid + 1
        } else {
            ans = mid
            r = mid - 1
        }
    }
    return ans
}

// useTraverse time complexity O(M+N), space complexity O(1)
func useTraverse(board [][]int) int {
    m := len(board)
    ans := -1
    if m < 1 {
        return ans
    }
    n := len(board[0])
    if n < 1 {
        return ans
    }
    var i, j int
    j = n - 1
    for i < m && j >= 0 {
        if board[i][j] == 1 {
            if ans != -1 {
                ans = utils.Min(ans, j)
            } else {
                ans = j
            }
            j--
        } else {
            i++
        }
    }
    return ans
}
