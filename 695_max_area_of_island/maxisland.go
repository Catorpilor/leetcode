package maxisland

func maxArea(grid [][]int) int {
    return dfs(grid)
}

// dfs time complexity is O(M*N), space complexity is O(1)
func dfs(grid [][]int) int {
    m := len(grid)
    if m < 1 {
        return 0
    }
    n := len(grid[0])
    if n < 1 {
        return 0
    }
    var res int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                cval := helper(grid, i, j, m, n)
                if cval > res {
                    res = cval
                }
            }
        }
    }
    return res
}

func helper(grid [][]int, i, j, m, n int) int {
    var res int
    if grid[i][j] == 1 {
        res += 1
        grid[i][j] = 0
    } else {
        return res
    }
    if i > 0 {
        res += helper(grid, i-1, j, m, n)
    }
    if i < m-1 {
        res += helper(grid, i+1, j, m, n)
    }
    if j > 0 {
        res += helper(grid, i, j-1, m, n)
    }
    if j < n-1 {
        res += helper(grid, i, j+1, m, n)
    }
    return res
}
