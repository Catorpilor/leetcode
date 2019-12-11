package sn

import (
    "reflect"
    "testing"
)

func TestCountingStepNumbers(t *testing.T) {
    st := []struct {
        name    string
        low, hi int
        exp     []int
    }{
        {"low hi eq 0", 0, 0, []int{0}},
        {"testcase1", 10, 12, []int{10, 12}},
        {"testcase2", 0, 21, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 21}},
        {"testcase3", 0, 8, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := countingStepNumbers(tt.low, tt.hi)
            if !reflect.DeepEqual(tt.exp, out) {
                t.Fatalf("with input low:%d and hi:%d wanted %v but got %v", tt.low, tt.hi, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
