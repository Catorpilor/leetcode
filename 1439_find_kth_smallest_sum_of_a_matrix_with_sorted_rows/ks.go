package ks

func kthSmallest(matrix [][]int, k int) int {
    return useBinarySearch(matrix, k)
}

// useBinarySearch time complexity O(log(5000*M)*M*K), space complexity O(M)
func useBinarySearch(matrix [][]int, k int) int {
    m := len(matrix)
    if m < 1 {
        return 0
    }
    n := len(matrix[0])
    if n < 1 {
        return 0
    }
    left, right, ans := m, 5000*m, -1
    for left <= right {
        mid := left + (right-left)/2
        cnt := countArraysHaveSumLessOrEqual(matrix, m, n, mid, 0, 0, k)
        if cnt >= k {
            right = mid - 1
            ans = mid
        } else {
            left = mid + 1
        }
    }
    return ans
}

// countArraysHaveSumLessOrEqual time complexity O(MK), space complexity O(M)
// countArraysHaveSumLessOrEqual() can run up to (m-i+1) * min(k,n^i), 1<=i<=m times.
// And n^i can go up to k very quickly while i is small, so time complexity will be O(m*k)
func countArraysHaveSumLessOrEqual(matrix [][]int, m, n, targetSum, row, sum, k int) int {
    if sum > targetSum {
        return 0
    }
    if row == m {
        return 1
    }
    var ans int
    // min(k, n^i)
    for c := 0; c < n; c++ {
        cnt := countArraysHaveSumLessOrEqual(matrix, m, n, targetSum, row+1, sum+matrix[row][c], k-ans)
        if cnt == 0 {
            break
        }
        ans += cnt
        if ans > k {
            break
        }
    }
    return ans
}
