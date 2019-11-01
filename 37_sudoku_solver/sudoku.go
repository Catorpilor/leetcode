package sudoku

const (
    n = 3
    N = 3 * 3
)

var (
    resolved bool
)

func solveSudoku(board [][]byte) [][]byte {
    // box size 3
    row := make([][]int, N)
    columns := make([][]int, N)
    box := make([][]int, N)
    res := make([][]byte, N)
    for i := 0; i < N; i++ {
        row[i] = make([]int, N+1)
        columns[i] = make([]int, N+1)
        box[i] = make([]int, N+1)
        res[i] = make([]byte, N)
        copy(res[i], board[i])
    }

    for i := 0; i < N; i++ {
        for j := 0; j < N; j++ {
            if board[i][j] != '.' {
                placeNumberAtPos(&res, i, j, int(board[i][j]-'0'), &row, &columns, &box)
            }
        }
    }
    permute(&res, 0, 0, &row, &columns, &box)
    return res
}

func placeNumberAtPos(board *[][]byte, i, j, num int, row, columns, box *[][]int) {
    boxIdx := (i/n)*n + j/n
    (*row)[i][num]++
    (*columns)[j][num]++
    (*box)[boxIdx][num]++
    (*board)[i][j] = byte('0' + num)
}

func removeNumberAtPos(board *[][]byte, i, j, num int, row, columns, box *[][]int) {
    boxIdx := (i/n)*n + j/n
    (*row)[i][num]++
    (*columns)[j][num]++
    (*box)[boxIdx][num]++
    (*board)[i][j] = '.'
}

func isValidPosForNum(i, j, num int, row, columns, box [][]int) bool {
    boxIdx := (i/n)*n + j/n
    return row[i][num]+columns[j][num]+box[boxIdx][num] == 0
}

func permute(board *[][]byte, i, j int, row, column, box *[][]int) {
    if (*board)[i][j] == '.' {
        for k := 1; k <= N; k++ {
            if isValidPosForNum(i, j, k, *row, *column, *box) {
                placeNumberAtPos(board, i, j, k, row, column, box)
                placeNext(board, i, j, row, column, box)
                if !resolved {
                    removeNumberAtPos(board, i, j, k, row, column, box)
                }
            }
        }
    } else {
        placeNext(board, i, j, row, column, box)
    }
}

func placeNext(board *[][]byte, i, j int, row, column, box *[][]int) {
    if i == N-1 && j == N-1 {
        resolved = true
    }
    if j == N-1 {
        permute(board, i+1, 0, row, column, box)
    } else {
        permute(board, i, j+1, row, column, box)
    }
}
