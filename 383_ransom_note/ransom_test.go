package ransom

import "testing"

func TestCanConstruct(t *testing.T) {
    st := []struct {
        name                 string
        ransomNote, magazine string
        exp                  bool
    }{
        {"both empty", "", "", true},
        {"magazine is empty", "abc", "", false},
        {"ransom is empty", "", "abc", true},
        {"testcase1", "aa", "aab", true},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := canConstruct(tt.ransomNote, tt.magazine)
            if out != tt.exp {
                t.Fatalf("with input ransom:%s and magazine:%s wanted %t but got %t", tt.ransomNote, tt.magazine, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
