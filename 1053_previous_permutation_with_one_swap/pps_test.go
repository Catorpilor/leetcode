package pps

import (
    "reflect"
    "testing"
)

func TestPrevPermOne(t *testing.T) {
    st := []struct {
        name string
        nums []int
        exp  []int
    }{
        {"single elemnt", []int{1}, []int{1}},
        {"all identical", []int{1, 1, 1}, []int{1, 1, 1}},
        {"decresing", []int{3, 2, 1}, []int{3, 1, 2}},
        {"testcase1", []int{1, 2, 3}, []int{1, 2, 3}},
        {"testcase2", []int{1, 9, 4, 6, 7}, []int{1, 7, 4, 6, 9}},
        {"testcase3", []int{3, 1, 1, 3}, []int{1, 3, 1, 3}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := prevPerm(tt.nums)
            if !reflect.DeepEqual(tt.exp, out) {
                t.Fatalf("with input nums:%v wanted %v but got %v", tt.nums, tt.exp, out)
            }
            t.Log("pass")
        })
    }

}
