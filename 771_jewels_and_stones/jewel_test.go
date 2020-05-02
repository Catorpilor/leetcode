package jewel

import "testing"

func TestNumOfStones(t *testing.T) {
    st := []struct {
        name string
        j, s string
        exp  int
    }{
        {"empty j", "", "abckd", 0},
        {"testcase1", "aA", "aAAbbbbb", 3},
        {"testcase2", "z", "ZZ", 0},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := numOfStones(tt.j, tt.s)
            if out != tt.exp {
                t.Fatalf("with input J:%s and S:%s wanted %d but got %d", tt.j, tt.s, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
