package ca

import "testing"

func TestMaxEvenCount(t *testing.T) {
    st := []struct {
        name string
        nums []int
        exp  int
    }{
        {"empty nums", nil, 0},
        {"single element", []int{1}, 0},
        {"all identical", []int{1, 1, 1, 1, 1}, 0},
        {"testcase1", []int{1, 1, 0, 1, 0, 0, 0, 1}, 8},
        {"testcase2", []int{1, 0, 1}, 2},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := maxContiguousArray(tt.nums)
            if out != tt.exp {
                t.Fatalf("with input nums: %v wanted %d but got %d", tt.nums, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestMaxEvenCountUseHashMap(t *testing.T) {
    st := []struct {
        name string
        nums []int
        exp  int
    }{
        {"empty nums", nil, 0},
        {"single element", []int{1}, 0},
        {"all identical", []int{1, 1, 1, 1, 1}, 0},
        {"testcase1", []int{1, 1, 0, 1, 0, 0, 0, 1}, 8},
        {"testcase2", []int{1, 0, 1}, 2},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := useHashmap(tt.nums)
            if out != tt.exp {
                t.Fatalf("with input nums: %v wanted %d but got %d", tt.nums, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
