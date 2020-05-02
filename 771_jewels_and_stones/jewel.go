package jewel

func numOfStones(j, s string) int {
    return useHashmap(j, s)
}

// useHashmap time complexity O(M+N), space complexity O(M)
func useHashmap(j, s string) int {
    n := len(j)
    if n < 1 {
        return n
    }
    hset := make(map[byte]bool, n)
    for i := range j {
        hset[j[i]] = true
    }
    var ans int
    for i := range s {
        if hset[s[i]] {
            ans++
        }
    }
    return ans
}
