package removedups

import "strings"

func removeDuplicates(s string) string {
    return useStack(s)
}

// useStack time complexity O(N), space complexity O(N)
func useStack(s string) string {
    n := len(s)
    if n <= 1 {
        return s
    }
    last_occ := make(map[byte]int, 26)
    for i := range s {
        last_occ[s[i]] = i
    }
    seen := make(map[byte]bool, 26)
    // stack stores the pos in s
    st := make([]int, 0, n)
    // iterate s we find the min s[i]
    for i := range s {
        if !seen[s[i]] {
            nst := len(st)
            // lexicographical order means "a" < "[b-z]*"
            // so we need to put the smallest s[i] in the begining.
            for nst > 0 && s[i] < s[st[nst-1]] && last_occ[s[st[nst-1]]] > i {
                // remove from seen
                delete(seen, s[st[nst-1]])
                // pop from stack
                st = st[:nst-1]
                nst--
            }
            seen[s[i]] = true
            st = append(st, i)
        }
    }
    var sb strings.Builder
    for i := range st {
        sb.WriteByte(s[st[i]])
    }
    return sb.String()
}
