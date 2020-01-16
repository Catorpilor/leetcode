package sm

import (
    "reflect"
    "testing"
)

func TestGenerateMatrix(t *testing.T) {
    st := []struct {
        name string
        n    int
        exp  [][]int
    }{
        {"n=1", 1, [][]int{[]int{1}}},
        {"n=2", 2, [][]int{[]int{1, 2}, []int{4, 3}}},
        {"n=3", 3, [][]int{[]int{1, 2, 3}, []int{8, 9, 4}, []int{7, 6, 5}}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := generateMatrix(tt.n)
            if !reflect.DeepEqual(tt.exp, out) {
                t.Fatalf("with n:%d wanted %v but got %v", tt.n, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
