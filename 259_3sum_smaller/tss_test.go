package tss

import "testing"

func Test3SumSmaller(t *testing.T) {
    st := []struct {
        name        string
        nums        []int
        target, exp int
    }{
        {"empty slice", []int{}, 1, 0},
        {"len(nums) < 3", []int{1, 2}, 3, 0},
        {"no matches", []int{1, 2, 3}, 4, 0},
        {"testcase1", []int{3, 0, 1, -2}, 2, 2},
        {"failed1", []int{1, -1, 2, 0, 3, -2}, 2, 10},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := sumSmaller(tt.nums, tt.target)
            if out != tt.exp {
                t.Fatalf("with nums:%v and target:%d wanted %d but got %d",
                    tt.nums, tt.target, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
