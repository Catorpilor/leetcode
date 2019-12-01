package ciw

// canIWin solve using dfs+backtrakcing.
// with memorization time complexity is o(2^n)
// without memorization time complexity is O(n!)
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
    if target <= 0 {
        return false
    }
    key := format(used)
    v, exists := set[key]
    if !exists {
        for i := 1; i < len(used); i++ {
            if !used[i] {
                used[i] = true
                // player #1 pick i, and player #2 lose
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
