package pawf

import "math"

func pacificAtlantic(matrix [][]int) [][]int {
    return dfs(matrix)
}

// dfs time complexity O(MN), space complexity O(MN)
func dfs(matrix [][]int) [][]int {
    var res [][]int
    m := len(matrix)
    if m < 1 {
        return res
    }
    n := len(matrix[0])
    if n < 1 {
        return res
    }
    pacific, atlantic := make([][]bool, m), make([][]bool, m)
    for i := 0; i < m; i++ {
        pacific[i] = make([]bool, n)
        atlantic[i] = make([]bool, n)
    }
    // start from the boarder
    for i := 0; i < m; i++ {
        // left boader with pacific ocean
        helper(matrix, pacific, math.MinInt32, i, 0, m, n)
        // right boarder with atlantic ocean
        helper(matrix, atlantic, math.MinInt32, i, n-1, m, n)
    }

    for j := 0; j < n; j++ {
        // top boarder with pacific ocean
        helper(matrix, pacific, math.MinInt32, 0, j, m, n)
        // bottom boarder with atlantic ocean
        helper(matrix, atlantic, math.MinInt32, m-1, j, m, n)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if pacific[i][j] && atlantic[i][j] {
                res = append(res, []int{i, j})
            }
        }
    }
    return res
}

func helper(matrix [][]int, visited [][]bool, height, x, y, m, n int) {
    if x < 0 || y < 0 || x >= m || y >= n || visited[x][y] || matrix[x][y] < height {
        return
    }
    visited[x][y] = true
    dirs := [5]int{-1, 0, 1, 0, -1}
    for i := 0; i < 4; i++ {
        nx, ny := x+dirs[i], y+dirs[i+1]
        helper(matrix, visited, matrix[x][y], nx, ny, m, n)
    }
}

// bfs time complexity O(MN), space complexity O(MN)
func bfs(matrix [][]int) [][]int {
    var res [][]int
    m := len(matrix)
    if m < 1 {
        return res
    }
    n := len(matrix[0])
    if n < 1 {
        return res
    }
    pacific, atlantic := make([][]bool, m), make([][]bool, m)
    for i := 0; i < m; i++ {
        pacific[i] = make([]bool, n)
        atlantic[i] = make([]bool, n)
    }
    type pair struct {
        x, y int
    }
    plevel, alevel := make([]pair, 0, m*n), make([]pair, 0, m*n)
    // start from the boarder
    for i := 0; i < m; i++ {
        // left boader with pacific ocean
        pacific[i][0] = true
        plevel = append(plevel, pair{i, 0})
        // right boarder with atlantic ocean
        atlantic[i][n-1] = true
        alevel = append(alevel, pair{i, n - 1})
    }

    for j := 0; j < n; j++ {
        // top boarder with pacific ocean
        pacific[0][j] = true
        plevel = append(plevel, pair{0, j})
        // bottom boarder with atlantic ocean
        atlantic[m-1][j] = true
        alevel = append(alevel, pair{m - 1, j})
    }
    dirs := [5]int{-1, 0, 1, 0, -1}
    for len(plevel) != 0 {
        next := make([]pair, 0, m*n)
        for _, p := range plevel {
            for i := 0; i < 4; i++ {
                nx, ny := p.x+dirs[i], p.y+dirs[i+1]
                if nx < 0 || nx >= m || ny < 0 || ny >= n || pacific[nx][ny] || matrix[nx][ny] < matrix[p.x][p.y] {
                    continue
                }
                pacific[nx][ny] = true
                next = append(next, pair{nx, ny})
            }
        }
        plevel = next
    }

    for len(alevel) != 0 {
        next := make([]pair, 0, m*n)
        for _, p := range alevel {
            for i := 0; i < 4; i++ {
                nx, ny := p.x+dirs[i], p.y+dirs[i+1]
                if nx < 0 || nx >= m || ny < 0 || ny >= n || atlantic[nx][ny] || matrix[nx][ny] < matrix[p.x][p.y] {
                    continue
                }
                atlantic[nx][ny] = true
                next = append(next, pair{nx, ny})
            }
        }
        alevel = next
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if pacific[i][j] && atlantic[i][j] {
                res = append(res, []int{i, j})
            }
        }
    }

    return res
}
