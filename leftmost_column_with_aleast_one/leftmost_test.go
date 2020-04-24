package leftmost

import "testing"

func TestLeftmostColumnWithOne(t *testing.T) {
    st := []struct {
        name  string
        board [][]int
        exp   int
    }{
        {"empty board", nil, -1},
        {"all zeors", [][]int{[]int{0, 0}, []int{0, 0}}, -1},
        {"all ones", [][]int{[]int{1, 1}, []int{1, 1}}, 0},
        {"testcase1", [][]int{[]int{0, 0}, []int{0, 1}}, 1},
        {"testcsae2", [][]int{[]int{0, 0, 0}, []int{0, 0, 1}, []int{0, 1, 1}}, 1},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := leftMostColumnWithOne(tt.board)
            if out != tt.exp {
                t.Fatalf("with input board:%v wanted %d but got %d", tt.board, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
