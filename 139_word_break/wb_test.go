package wb

import "testing"

func TestWordBreak(t *testing.T) {
	st := []struct {
		name, s string
		word    []string
		exp     bool
	}{
		{"empty str", "", []string{"test"}, false},
		{"testcase1", "catdog", []string{"cat", "dog"}, true},
		{"testcase2", "applepenapple", []string{"apple", "pen"}, true},
		{"testcase3", "catsandog", []string{"cats", "cat", "sand", "dog", "and"}, false},
		{"failed1", "cars", []string{"car", "ca", "rs"}, true},
		{"failed2", "bb", []string{"a", "b", "bb", "bbbb"}, true},
		// tle
		{"tle", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}, false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := wordBreak(tt.s, tt.word)
			if out != tt.exp {
				t.Fatalf("with input s:%s and words: %v wanted %t but got %t", tt.s, tt.word, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
