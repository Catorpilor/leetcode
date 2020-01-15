package gol

import (
    "reflect"
    "testing"
)

func TestGameOfLife(t *testing.T) {
    st := []struct {
        name  string
        board [][]int
        exp   [][]int
    }{
        {"board is empty", [][]int{}, [][]int{}},
        {"testcase1", [][]int{[]int{0, 1, 0}, []int{0, 0, 1}, []int{1, 1, 1}, []int{0, 0, 0}},
            [][]int{[]int{0, 0, 0}, []int{1, 0, 1}, []int{0, 1, 1}, []int{0, 1, 0}}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := gameOfLife(tt.board)
            if !reflect.DeepEqual(tt.exp, out) {
                t.Fatalf("with board: %v wanted %v but got %v", tt.board, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestOptimized(t *testing.T) {
    st := []struct {
        name  string
        board [][]int
        exp   [][]int
    }{
        {"board is empty", [][]int{}, [][]int{}},
        {"testcase1", [][]int{[]int{0, 1, 0}, []int{0, 0, 1}, []int{1, 1, 1}, []int{0, 0, 0}},
            [][]int{[]int{0, 0, 0}, []int{1, 0, 1}, []int{0, 1, 1}, []int{0, 1, 0}}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := spaceOptimized(tt.board)
            if !reflect.DeepEqual(tt.exp, out) {
                t.Fatalf("with board: %v wanted %v but got %v", tt.board, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
