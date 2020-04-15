package removedups

import "testing"

func TestRemoveDups(t *testing.T) {
    st := []struct {
        name string
        s    string
        exp  string
    }{
        {"empty string", "", ""},
        {"all identical", "cccc", "c"},
        {"testcase1", "bcabc", "abc"},
        {"testcase2", "cbacdcbc", "acdb"},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := removeDuplicates(tt.s)
            if out != tt.exp {
                t.Fatalf("with input s:%s wanted %s but got %s", tt.s, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
