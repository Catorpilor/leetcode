package bp

import "testing"

func TestBreakPalindrome(t *testing.T) {
	st := []struct {
		name, s, exp string
	}{
		{"single character", "a", ""},
		{"aa", "aa", "ab"},
		{"aba", "aba", "abb"},
		{"testcase1", "abccba", "aaccba"},
		{"testcase2", "aacaa", "aacab"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := breakPalindrome(tt.s)
			if out != tt.exp {
				t.Fatalf("with s:%s wanted %s but got %s", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
