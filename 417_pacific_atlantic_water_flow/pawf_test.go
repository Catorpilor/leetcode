package pawf

import (
    "reflect"
    "testing"
)

func TestWaterFlow(t *testing.T) {
    st := []struct {
        name    string
        maxtrix [][]int
        exp     [][]int
    }{
        {"empty matrix", [][]int{}, [][]int{}},
        {"testcase1", [][]int{[]int{1, 2, 2, 3, 5}, []int{3, 2, 3, 4, 4}, []int{2, 4, 5, 3, 1}, []int{6, 7, 1, 4, 5}, []int{5, 1, 1, 2, 4}},
            [][]int{[]int{0, 4}, []int{1, 3}, []int{1, 4}, []int{2, 2}, []int{3, 0}, []int{3, 1}, []int{4, 0}}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := pacificAtlantic(tt.maxtrix)
            if !reflect.DeepEqual(tt.exp, out) {
                t.Fatalf("with matrix: %v wanted %v but got %v", tt.maxtrix, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestWaterFlowBfs(t *testing.T) {
    st := []struct {
        name    string
        maxtrix [][]int
        exp     [][]int
    }{
        {"empty matrix", [][]int{}, [][]int{}},
        {"testcase1", [][]int{[]int{1, 2, 2, 3, 5}, []int{3, 2, 3, 4, 4}, []int{2, 4, 5, 3, 1}, []int{6, 7, 1, 4, 5}, []int{5, 1, 1, 2, 4}},
            [][]int{[]int{0, 4}, []int{1, 3}, []int{1, 4}, []int{2, 2}, []int{3, 0}, []int{3, 1}, []int{4, 0}}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := bfs(tt.maxtrix)
            if !reflect.DeepEqual(tt.exp, out) {
                t.Fatalf("with matrix: %v wanted %v but got %v", tt.maxtrix, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
