package substr

import "testing"

func TestLongestSubstr(t *testing.T) {
    st := []struct {
        name string
        s    string
        exp  int
    }{
        {"empty string", "", 0},
        {"single character", "a", 1},
        {"all identical", "aaa", 1},
        {"testcase1", "abcabc", 3},
        {"testcase2", "pwwkew", 3},
        {"testcase4", "abc122 cddf", 5},
        {"failed1", "wobgrovw", 6},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := longestSubstr(tt.s)
            if out != tt.exp {
                t.Fatalf("with input s:%s wanted %d but got %d", tt.s, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
