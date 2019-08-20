package maze

import "testing"

func TestCherryPickUp(t *testing.T) {
    st := []struct {
        name  string
        input [][]int
        exp   int
    }{
        {"N=1", [][]int{[]int{1}}, 1},
        {"N=1 without cherry", [][]int{[]int{0}}, 0},
        {"N=2 filled with cherries", [][]int{[]int{1, 1}, []int{1, 1}}, 4},
        {"N=2, no path", [][]int{[]int{1, -1}, []int{-1, 1}}, 0},
        {"n=3 testcase 1", [][]int{[]int{0, 1, -1}, []int{1, 0, -1}, []int{1, 1, 1}}, 5},
        {"n=7 failed", [][]int{[]int{1, 1, 1, 1, 0, 0, 0}, []int{0, 0, 0, 1, 0, 0, 0},
            []int{0, 0, 0, 1, 0, 0, 1}, []int{1, 0, 0, 1, 0, 0, 0}, []int{0, 0, 0, 1, 0, 0, 0},
            []int{0, 0, 0, 1, 0, 0, 0}, []int{0, 0, 0, 1, 1, 1, 1}}, 15},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            ret := CherryPickUp(tt.input)
            if ret != tt.exp {
                t.Fatalf("with input %v wanted %d but got %d", tt.input, tt.exp, ret)
            }
            t.Log("pass")
        })
    }
}
