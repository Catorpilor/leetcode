package gol

func gameOfLife(board [][]int) [][]int {
    return bruteForce(board)
}

type pos struct {
    x, y int
}

// bruteForce time compelxity is O(N*M), space complexy is O(N*M)
func bruteForce(board [][]int) [][]int {
    n := len(board)
    if n < 1 {
        return board
    }
    m := len(board[0])
    if m < 1 {
        return board
    }
    bak := make([][]int, n)
    for i := range bak {
        bak[i] = make([]int, m)
        for j := range bak[i] {
            bak[i][j] = board[i][j]
        }
    }

    hset := make(map[pos]int, n*m)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            numOfLives := process(bak, i, j, n, m)
            if bak[i][j] == 1 {
                if numOfLives < 2 || numOfLives > 3 {
                    hset[pos{i, j}] = 0
                }
            } else {
                if numOfLives == 3 {
                    hset[pos{i, j}] = 1
                }
            }
        }
    }
    for k, v := range hset {
        bak[k.x][k.y] = v
    }
    return bak
}
func process(board [][]int, i, j, n, m int) int {
    var res int
    if i > 0 {
        res += board[i-1][j]
    }
    if i < n-1 {
        res += board[i+1][j]
    }
    if j > 0 {
        res += board[i][j-1]
    }
    if j < m-1 {
        res += board[i][j+1]
    }
    if i > 0 && j > 0 {
        res += board[i-1][j-1]
    }
    if i > 0 && j < m-1 {
        res += board[i-1][j+1]
    }
    if i < n-1 && j > 0 {
        res += board[i+1][j-1]
    }
    if i < n-1 && j < m-1 {
        res += board[i+1][j+1]
    }
    return res
}
