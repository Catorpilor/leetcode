package nos

import "github.com/catorpilor/leetcode/utils"

func maxSum(A []int, l, m int) int {
    n := len(A)
    if n < 1 {
        return 0
    }
    presum := make([]int, n)
    presum[0] = A[0]
    for i := 1; i < n; i++ {
        presum[i] = presum[i-1] + A[i]
    }
    // return withPresum(presum, l, m)
    return utils.Max(cal(presum, l, m), cal(presum, m, l))
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
