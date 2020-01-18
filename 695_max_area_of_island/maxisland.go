package maxisland

import "github.com/catorpilor/leetcode/utils/uf"

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

func withUnionFind(grid [][]int) int {
    m := len(grid)
    if m < 1 {
        return 0
    }
    n := len(grid[0])
    if n < 1 {
        return 0
    }
    wqc := uf.NewWQUPC(m * n)
    var idx int
    dirs := [4][2]int{[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0}}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                idx = i*n + j
                for _, v := range dirs {
                    ni, nj := i+v[0], i+v[1]
                    nidx := ni*n + nj
                    if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == 1 {
                        wqc.Union(idx, nidx)
                    }
                }
            }
        }
    }
    hset := make(map[int]int, n*m)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                idx = i*n + j
                root := wqc.Root(idx)
                if _, exists := hset[root]; !exists {
                    hset[root] = 1
                } else {
                    hset[root]++
                }
            }
        }
    }
    var res int
    for _, v := range hset {
        if v > res {
            res = v
        }
    }
    return res
}
