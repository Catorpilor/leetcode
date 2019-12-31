package nos

import "github.com/catorpilor/leetcode/utils"

func maxSum(A []int, l, m int) int {
    n := len(A)
    if n < 1 {
        return 0
    }
    presum := make([]int, n+1)
    for i := 1; i <= n; i++ {
        presum[i] = presum[i-1] + A[i-1]
    }
    // return withPresum(presum, l, m)
    // return utils.Max(cal(presum, l, m), cal(presum, m, l))
    return utils.Max(burteForce(presum, l, m), burteForce(presum, m, l))
}

func withPresum(preSum []int, l, m int) int {
    res, lmax, mmax := preSum[l+m-1], preSum[l-1], preSum[m-1]
    for i := l + m; i < len(preSum); i++ {
        // lmax represents the max L subarray before M
        lmax = utils.Max(lmax, preSum[i-m]-preSum[i-m-l])
        // mmax represents the max M subarray before L
        mmax = utils.Max(mmax, preSum[i-l]-preSum[i-l-m])
        // lmax+presum[i]-presum[i-m] represents L subarray before M subarray
        // mmax+presum[i]-presum[i-l] represents M subarray before L subarray
        res = utils.Max(res, utils.Max(lmax+preSum[i]-preSum[i-m], mmax+preSum[i]-preSum[i-l]))
    }
    return res
}

func cal(preSum []int, l, m int) int {
    res, lmax := preSum[l+m-1], preSum[l-1]
    for i := l + m; i < len(preSum); i++ {
        // lmax stores the max L continous subarray sum that before M
        lmax = utils.Max(lmax, preSum[i-m]-preSum[i-m-l])
        res = utils.Max(res, lmax+preSum[i]-preSum[i-m])
    }
    return res
}

// bruteForce with presum time complexity O(n^2), Space Complexity O(n)
func burteForce(presum []int, l, m int) int {
    var res int
    n := len(presum)
    for i := 0; i < n-1-m; i++ {
        for j := i + l; j+m-1 < n-1; j++ {
            suml := presum[i+l] - presum[i]
            summ := presum[j+m] - presum[j]
            res = utils.Max(res, suml+summ)
        }
    }
    return res
}
