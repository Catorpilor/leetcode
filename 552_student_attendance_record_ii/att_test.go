package att

import "testing"

func TestCheckRecord(t *testing.T) {
    st := []struct {
        name string
        n    int
        exp  int
    }{
        {"edge 0", 0, 0},
        {"eq 1", 1, 3},
        {"eq 2", 2, 8},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := CheckRecord(tt.n)
            if out != tt.exp {
                t.Fatalf("with input n: %d wanted %d but got %d", tt.n,
                    tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
