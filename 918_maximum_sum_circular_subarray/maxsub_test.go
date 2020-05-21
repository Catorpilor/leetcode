package maxsub

import "testing"

func TestMaxSubSum(t *testing.T) {
    st := []struct {
        name string
        nums []int
        exp  int
    }{
        {"empty", nil, 0},
        {"single element", []int{1}, 1},
        {"all positive", []int{2, 3, 4}, 9},
        {"all neg", []int{-2, -3, -1}, -1},
        {"testcase1", []int{5, -3, 5}, 10},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := maxSubSum(tt.nums)
            if out != tt.exp {
                t.Fatalf("with input nums:%v wanted %d but got %d", tt.nums, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestUseKadane(t *testing.T) {
    st := []struct {
        name string
        nums []int
        exp  int
    }{
        {"all negs", []int{-2, -3, -1}, -1},
        {"all pos", []int{1, 2, 3}, 6},
        {"testcase1", []int{5, -3, 5}, 7},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := useKadane(tt.nums, len(tt.nums))
            if out != tt.exp {
                t.Fatalf("with input nums: %v wanted %d but got %d", tt.nums, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
