package wb

func wordBreak(s string, words []string) []string {
    n := len(s)
    nw := len(words)
    var res []string
    if n == 0 || nw == 0 {
        return res
    }
    set := make(map[string]bool, nw)
    for _, word := range words {
        set[word] = true
    }
    cache := make(map[int]bool, n)
    permute(&res, s, set, cache, "", 0, n)
    return res
}

func permute(res *[]string, s string, set map[string]bool, cache map[int]bool, bid string, idx, n int) {
    if idx == n {
        *res = append(*res, bid)
        return
    }
    if cache[idx] {
        return
    }
    for i := idx + 1; i <= n; i++ {
        if set[s[idx:i]] {
            pre := len(*res)
            temp := s[idx:i]
            if bid != "" {
                temp = bid + " " + s[idx:i]
            }
            permute(res, s, set, cache, temp, i, n)
            if len(*res) == pre {
                cache[i] = true
            }
        }
    }
    cache[idx] = true
    return
}
