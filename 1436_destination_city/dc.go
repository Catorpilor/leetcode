package dc

func destCity(paths [][]string) string {
    return useHashmap(paths)
}

func useHashmap(paths [][]string) string {
    n := len(paths)
    if n < 1 {
        return ""
    }
    in, out := make(map[string]int, n), make(map[string]int, n)
    for i := range paths {
        out[paths[i][0]]++
        in[paths[i][1]]++
    }
    var ans string
    for k := range in {
        if _, exists := out[k]; exists {
            continue
        }
        ans = k
        break
    }
    return ans
}
