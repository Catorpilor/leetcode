package spm

func spiralOrder(matrix [][]int) []int {
    return simulation(matrix)
}

// simulation time complexity is O(m*n), space complexity is O(m*n)
func simulation(matrix [][]int) []int {
    m := len(matrix)
    if m < 1 {
        return []int{}
    }
    n := len(matrix[0])
    if n < 1 {
        return []int{}
    }
    res := make([]int, 0, m*n)
    // seen stores the visited cells
    seen := make([]bool, m*n)
    // clockwise,
    // start at(i,j),so the next pos is (i,j+1) .. (i,n-1)
    // turn down iterator rows, the next pos is (i+1, n-1)..
    dr := [4]int{0, 1, 0, -1}
    dc := [4]int{1, 0, -1, 0}
    var r, c, di int
    for i := 0; i < m*n; i++ {
        res = append(res, matrix[r][c])
        seen[r*n+c] = true
        nr := r + dr[di]
        nc := c + dc[di]
        if (nr < m && nr >= 0) && (nc < n && nc >= 0) && !seen[nr*n+nc] {
            // update r, c
            r, c = nr, nc
        } else {
            // reach boundary, update di
            di = (di + 1) % 4
            r += dr[di]
            c += dc[di]
        }
    }
    return res
}
