package number

import "github.com/catorpilor/leetcode/utils"

func reachNumber(target int) int {
    return useMath(target)
}

func useMath(target int) int {
    target = utils.Abs(target)
    ans := 0
    var sum int
    for sum < target || (sum-target)%2 != 0 {
        ans++
        sum += ans
    }
    return ans
}

// useBruteForce got tle
func useBruteForce(target int) int {
    store := make(map[int]map[int]bool)
    store[1] = map[int]bool{1: true, -1: true}
    store[2] = map[int]bool{3: true, 1: true, -1: true, -3: true}
    ans := 1
    for {
        if store[ans] == nil {
            pvn := len(store[ans-1])
            store[ans] = make(map[int]bool, pvn<<1)
            for k := range store[ans-1] {
                store[ans][k+ans] = true
                store[ans][k-ans] = true
                if target == k+ans || target == k-ans {
                    return ans
                }
            }
        } else {
            for k := range store[ans] {
                if k == target {
                    return ans
                }
            }
        }
        ans++
    }
}
