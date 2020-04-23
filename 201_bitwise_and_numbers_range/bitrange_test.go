package bitrange

import "testing"

func TestRangeBitwiseAnd(t *testing.T) {
    st := []struct {
        name string
        m, n int
        exp  int
    }{
        {"all zeros", 0, 0, 0},
        {"m=n", 5, 5, 5},
        {"testcase1", 4, 5, 4},
        {"testcase2", 0, 1, 0},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := rangeBitwiseAnd(tt.m, tt.n)
            if out != tt.exp {
                t.Fatalf("with input m:%d, n:%d wanted %d but got %d", tt.m, tt.n, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
