package ciw

import "testing"

func TestCanIWin(t *testing.T) {
    st := []struct {
        name                  string
        maxChoosable, desired int
        exp                   bool
    }{
        {"desired less or eq maxChoosable", 20, 10, true},
        {"desired larger than maxChoosable", 10, 11, false},
        {"testcase1", 10, 56, false},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := canIWin(tt.maxChoosable, tt.desired)
            if out != tt.exp {
                t.Fatalf("with input maxChoosable:%d and desired: %d wanted %t but got %t",
                    tt.maxChoosable, tt.desired, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
