package sn

import (
    "sort"
    "strconv"
)

func countingStepNumbers(low, hi int) []int {
    // return bruteForce(low, hi)
    var res []int
    for i := 0; i <= 9; i++ {
        dfs(low, hi, int64(i), &res)
    }
    sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
    return res
}

func bruteForce(low, hi int) []int {
    res := make([]int, 0, hi-low)
    if low == hi {
        res = append(res, low)
        return res
    }
    res = append(res, low)
    for i := low + 1; i < hi; i++ {
        if isStepping(i) {
            res = append(res, i)
        }
    }
    res = append(res, hi)
    return res
}

func isStepping(i int) bool {
    sb := strconv.Itoa(i)
    var prev, cur int
    for i := 0; i < len(sb)-1; i++ {
        //fmt.Printf("sb[i] is %d, sb[i+1] is %d\n", int(sb[i]-'0'), int(sb[i+1]-'0'))
        prev, cur = int(sb[i]-'0'), int(sb[i+1]-'0')
        if prev-cur != -1 && prev-cur != 1 {
            return false
        }
    }
    return true
}

// dfs version
// for single digit 0,1,2,3,4,5,6,7,8,9 all fit
// For number from 1 to 9, append a new digit to the end that is 1 away from current.
// For example, 1 will be 12 and 10, same for others.
// If new numbers are in range, add them to list.
func dfs(low, hi int, cur int64, res *[]int) {
    if cur >= int64(low) && cur <= int64(hi) {
        *res = append(*res, int(cur))
    }
    if cur == int64(0) || cur > int64(hi) {
        return
    }
    // Recursively, find the last digit of a number, append a new digit that is 1 away from current last digit.
    lastDigit := cur % 10
    inc, dec := cur*10+1+lastDigit, cur*10+lastDigit-1
    if lastDigit == 0 {
        dfs(low, hi, inc, res)
    } else if lastDigit == 9 {
        dfs(low, hi, dec, res)
    } else {
        dfs(low, hi, inc, res)
        dfs(low, hi, dec, res)
    }
}
