package square

func countSquare(matrix [][]int) int {
    return 0
}

// useBottomUpDP time complexity O(MN), space complexity O(MN)
func useBottomUpDP(matrix [][]int) int {
    m := len(matrix)
    if m < 1 {
        return 0
    }
    n := len(matrix[0])
    if n < 1 {
        return 0
    }
    var ans int
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
        dp[i][n-1] = matrix[i][n-1]
        ans += dp[i][n-1]
    }
    for j := 0; j < n-1; j++ {
        dp[m-1][j] = matrix[m-1][j]
        ans += dp[m-1][j]
    }
    // fmt.Println(m,n)
    for i := m - 2; i >= 0; i-- {
        for j := n - 2; j >= 0; j-- {
            if matrix[i][j] == 1 {
                // fmt.Println(i,j)
                dp[i][j] = min(dp[i+1][j], min(dp[i+1][j+1], dp[i][j+1])) + 1
                ans += dp[i][j]
            }
        }
    }
    // fmt.Println(dp)
    return ans
}
