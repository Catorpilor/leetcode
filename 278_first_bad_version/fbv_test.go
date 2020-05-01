package fbv

import "testing"

func TestFirstBadVersion(t *testing.T) {
    st := []struct {
        name string
        vers []int
        exp  int
    }{
        {"single ver", []int{0}, 0},
        {"testcase1", []int{1, 1, 1, 1, 0, 0, 0}, 4},
        {"testcase2", []int{1, 1, 0, 0, 0, 0}, 2},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := firstBadVersion(tt.vers)
            if out != tt.exp {
                t.Fatalf("with input vers: %v wanted %d but got %d", tt.vers, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
