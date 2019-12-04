package wp

func isMatch(pattern, str string) bool {
    cache := make(map[byte]string, len(pattern))
    set := make(map[string]bool)
    return helper(pattern, 0, str, 0, cache, set)
}

func helper(pattern string, pi int, str string, sj int, cache map[byte]string, set map[string]bool) bool {
    // fmt.Printf("pi: %d, sj: %d, cache: %v\n", pi, sj, cache)
    if pi == len(pattern) && sj == len(str) {
        return true
    }
    if pi == len(pattern) || sj == len(str) {
        return false
    }
    c := pattern[pi]
    if s, exists := cache[c]; exists {
        l := sj + len(s)
        if l > len(str) {
            l = len(str)
        }
        if str[sj:l] != s {
            return false
        }
        return helper(pattern, pi+1, str, sj+len(s), cache, set)
    }
    // backtracking here
    // for example pattern="abab" str="redblueredblue"
    // c = 'a', p = 'r'
    // c = 'b', p = 'e'
    // c = 'a', p = 'd' != 'r' so we backtrack
    // c = 'b', p = 'ed'
    // so on and so forth.
    for k := sj; k < len(str); k++ {
        p := str[sj : k+1]
        if set[p] {
            continue
        }
        cache[c] = p
        set[p] = true
        if helper(pattern, pi+1, str, k+1, cache, set) {
            return true
        }
        delete(cache, c)
        set[p] = false
    }
    return false
}
