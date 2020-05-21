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

    // more comments
    // for example with matrix:
    // [  0,1,2
    // 0 [1,1,1]
    // 1 [1,1,1]
    // 2 [1,1,1]
    // 3 [1,1,1]
    // ]
    // we just calcute the top left corner, which means at pos(i,j) we can
    // only go right or bottom to remove the duplicates.
    // after initialize dp
    // [  0,1,2
    // 0 [0,0,1]
    // 1 [0,0,1]
    // 2 [0,0,1]
    // 3 [1,1,1]
    // ]
    // after traverse the matrix, we got
    // [  0,1,2
    // 0 [3,2,1]
    // 1 [3,2,1]
    // 2 [2,2,1]
    // 3 [1,1,1]
    // ]
    // so at pos(0,0) we got 3 squares, [0,0] length=1,  length=2 and  length=3
    //  at pos(1,1) we got 2 squres, [1,1] length=1 and length=2
    // so the final answer just sum them up.
    // fmt.Println(dp)
    return ans
}
