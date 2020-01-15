package spm

import (
    "reflect"
    "testing"
)

func TestSpiralOrder(t *testing.T) {
    st := []struct {
        name   string
        matrix [][]int
        exp    []int
    }{
        {"empty matrix", [][]int{}, []int{}},
        {"1*m matrix", [][]int{[]int{3, 2, 4, 5}}, []int{3, 2, 4, 5}},
        {"n*1 matrix", [][]int{[]int{1}, []int{2}, []int{3}}, []int{1, 2, 3}},
        {"testcase1", [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := spiralOrder(tt.matrix)
            if !reflect.DeepEqual(out, tt.exp) {
                t.Fatalf("with matrix:%v wanted %v but got %v", tt.matrix, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
