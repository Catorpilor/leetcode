package ciw

func canIWin(maxChooseable, desired int) bool {
    if maxChooseable >= desired {
        return true
    }
    sum := (1 + maxChooseable) * maxChooseable / 2
    if sum < desired {
        return false
    }
    set := make(map[int]bool)
    used := make([]bool, maxChooseable+1)

    return helper(set, used, desired)
}

func helper(set map[int]bool, used []bool, target int) bool {
    // fmt.Printf("turn:%d, s:%d, e:%d, target:%d\n", turn, s, e, target)
    if target <= 0 {
        return false
    }
    key := format(used)
    v, exists := set[key]
    if !exists {
        for i := 1; i < len(used); i++ {
            if !used[i] {
                used[i] = true
                if !helper(set, used, target-i) {
                    set[key] = true
                    used[i] = false
                    return true
                }
                used[i] = false
            }
        }
        set[key] = false
    }
    return v

}

func format(used []bool) int {
    var res int
    for i := range used {
        res <<= 1
        if used[i] {
            res |= 1
        }
    }
    return res
}
