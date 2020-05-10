package judge

func findJudge(n int, trust [][]int) int {
    return useHashMap(n, trust)
}

// useHashMap time complexity O(N), space complexity O(N)
func useHashMap(n int, trust [][]int) int {
    if n == 1 {
        return 1
    }
    in, out := make(map[int]int, n), make(map[int]int, n)
    for i := range trust {
        in[trust[i][1]]++
        out[trust[i][0]]++
    }
    for k := range in {
        if in[k] == n-1 && out[k] == 0 {
            return k
        }
    }
    return -1
}
