package ip

func islandPerimeter(grid [][]int) int {
    return countWithEqu(grid)
}

// countWithEqu time complexity O(MN), space complexity O(1)
func countWithEqu(grid [][]int) int {
    m := len(grid)
    if m < 1 {
        return 0
    }
    n := len(grid[0])
    if n < 1 {
        return 0
    }
    var islands, neighbors int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                islands++
                if j < n-1 && grid[i][j+1] == 1 {
                    // right neighbors
                    neighbors++
                }
                if i < m-1 && grid[i+1][j] == 1 {
                    // down neighbors
                    neighbors++
                }
            }
        }
    }
    // fmt.Printf("res: %d\n", res)
    return islands*4 - neighbors*2
}

// iterator count the num of edges of island, time complexity O(MN)
func iterator(grid [][]int) int {
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
            // if grid[i][j] == 1 we were supposed to add 4 edges, but we've to check its neighbors
            if grid[i][j] == 1 {
                if i == 0 || grid[i-1][j] == 0 {
                    res++
                }
                if i == m-1 || grid[i+1][j] == 0 {
                    res++
                }
                if j == 0 || grid[i][j-1] == 0 {
                    res++
                }
                if j == n-1 || grid[i][j+1] == 0 {
                    res++
                }
            }
        }
    }
    return res
}
