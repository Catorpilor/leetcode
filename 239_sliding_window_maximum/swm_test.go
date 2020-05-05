package swm

import (
    "reflect"
    "testing"
)

func TestMaxSlidingWindow(t *testing.T) {
    st := []struct {
        name string
        nums []int
        k    int
        exp  []int
    }{
        {"nil slice", nil, 1, nil},
        {"single slice", []int{1}, 1, []int{1}},
        {"identical", []int{1, 1, 1, 1}, 2, []int{1, 1, 1}},
        {"testcase1", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := maxSlidingWindow(tt.nums, tt.k)
            if !reflect.DeepEqual(out, tt.exp) {
                t.Fatalf("with input nums:%v and k:%d, wanted %v but got %v", tt.nums, tt.k, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestMaxSlidingWindowUseDeque(t *testing.T) {
    st := []struct {
        name string
        nums []int
        k    int
        exp  []int
    }{
        {"nil slice", nil, 1, nil},
        {"single slice", []int{1}, 1, []int{1}},
        {"identical", []int{1, 1, 1, 1}, 2, []int{1, 1, 1}},
        {"testcase1", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := useDeque(tt.nums, tt.k)
            if !reflect.DeepEqual(out, tt.exp) {
                t.Fatalf("with input nums:%v and k:%d, wanted %v but got %v", tt.nums, tt.k, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
