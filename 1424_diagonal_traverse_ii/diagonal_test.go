package diagonal

import (
    "reflect"
    "testing"
)

func TestDiagonalOrder(t *testing.T) {
    st := []struct {
        name string
        nums [][]int
        exp  []int
    }{
        {"empty nums", nil, nil},
        {"one row", [][]int{[]int{1, 2, 3, 4, 5}}, []int{1, 2, 3, 4, 5}},
        {"testcase1", [][]int{[]int{1, 2, 3, 4, 5}, []int{6, 7}, []int{8}, []int{9, 10, 11}, []int{12, 13, 14, 15, 16}}, []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := diagonalTraverse(tt.nums)
            if !reflect.DeepEqual(out, tt.exp) {
                t.Fatalf("with input nums:%v wanted %v but got %v", tt.nums, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
