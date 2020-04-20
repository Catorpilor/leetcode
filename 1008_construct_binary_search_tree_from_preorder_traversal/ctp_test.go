package ctp

import (
    "math"
    "reflect"
    "testing"

    "github.com/catorpilor/leetcode/utils"
)

func TestConstructTree(t *testing.T) {
    st := []struct {
        name  string
        order []int
        exp   []int
    }{
        {"empty order", nil, nil},
        {"single element", []int{1}, []int{1}},
        {"testcase1", []int{8, 5, 1, 7, 10, 12}, []int{8, 5, 10, 1, 7, math.MinInt32, 12}},
        {"testcase2", []int{5, 3, 2, 4}, []int{5, 3, math.MinInt32, 2, 4}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            root := constructTree(tt.order)
            out := utils.LevelOrderTravesal(root)
            if !reflect.DeepEqual(tt.exp, out) {
                t.Fatalf("with input order: %v wanted %v but got %v", tt.order, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
