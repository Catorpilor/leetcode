package lcp

import (
    "reflect"
    "testing"
)

func TestLetterCombinations(t *testing.T) {
    st := []struct {
        name   string
        digits string
        exp    []string
    }{
        {"empty", "", []string{}},
        {"digits 22", "22", []string{"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"}},
        {"with 1", "123456789", []string{}},
        {"with 0", "2340", []string{}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := letterCombinations(tt.digits)
            if !reflect.DeepEqual(tt.exp, out) {
                t.Fatalf("with input digits: %s wanted %v but got %v", tt.digits, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
