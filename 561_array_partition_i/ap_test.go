package ap

import "testing"

func TestArrayPairSum(t *testing.T) {
    st := []struct {
        name string
        nums []int
        exp  int
    }{
        {"empty slice", []int{}, 0},
        {"odds slice", []int{1, 2, 3}, 0},
        {"all identical", []int{2, 2, 2, 2}, 4},
        {"testcase1", []int{1, 4, 2, 3}, 4},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := arrayPairSum(tt.nums)
            if out != tt.exp {
                t.Fatalf("with nums: %v wanted %d but got %d", tt.nums, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
