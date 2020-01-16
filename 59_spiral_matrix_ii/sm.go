package sm

func generateMatrix(n int) [][]int {
    return simulate(n)
}

// simulate time complexity O(N*N), space complexity O(N*N)
func simulate(n int) [][]int {
    dr, dc := [4]int{0, 1, 0, -1}, [4]int{1, 0, -1, 0}
    res := make([][]int, n)
    for i := range res {
        res[i] = make([]int, n)
    }
    seen := make([]bool, n*n)
    var r, c, di int
    for i := 1; i <= n*n; i++ {
        res[r][c] = i
        seen[r*n+c] = true
        nr := r + dr[di]
        nc := c + dc[di]
        if (nr >= 0 && nr < n) && (nc < n && nc >= 0) && !seen[nr*n+nc] {
            r, c = nr, nc
        } else {
            di = (di + 1) % 4
            r += dr[di]
            c += dc[di]
        }
    }
    return res
}
