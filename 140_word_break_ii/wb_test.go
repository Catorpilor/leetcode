package wb

import (
    "reflect"
    "testing"
)

func TestWordBreak(t *testing.T) {
    st := []struct {
        name, s string
        word    []string
        exp     []string
    }{
        {"empty str", "", []string{"test"}, []string{}},
        {"testcase1", "catdog", []string{"cat", "dog"}, []string{"cat dog"}},
        {"testcase2", "applepenapple", []string{"apple", "pen"}, []string{"apple pen apple"}},
        {"testcase3", "catsandog", []string{"cats", "cat", "sand", "dog", "and"}, []string{}},
        {"failed1", "cars", []string{"car", "ca", "rs"}, []string{"ca rs"}},
        {"failed2", "bb", []string{"a", "b", "bb", "bbbb"}, []string{"b b", "bb"}},
        // tle
        {"tle", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
            []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}, []string{}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := wordBreak(tt.s, tt.word)
            if !reflect.DeepEqual(tt.exp, out) {
                t.Fatalf("with input s:%s and words: %v wanted %v but got %v", tt.s, tt.word, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
