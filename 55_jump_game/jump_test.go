package jump

import "testing"

func TestCanJump(t *testing.T) {
    st := []struct {
        name string
        arr  []int
        exp  bool
    }{
        {"nil slice", nil, false},
        {"single one", []int{0}, true},
        {"more than one but starts with 0", []int{0, 1, 2, 3}, false},
        {"testcase1", []int{2, 3, 1, 1, 4}, true},
        {"testcase2", []int{2, 0, 2, 0, 0}, true},
    }
    for _, tt := range st {
        out := canJump(tt.arr)
        if out != tt.exp {
            t.Fatalf("with input arr:%v wanted %t but got %t", tt.arr, tt.exp, out)
        }
        t.Log("pass")
    }
}
