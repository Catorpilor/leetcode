package sn

import "strconv"

func countingStepNumbers(low, hi int) []int {
    return bruteForce(low, hi)
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
