package palin

import "testing"

func TestPalin(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  bool
	}{
		{"empty string", "", true},
		{"single string", "a", true},
		{"len eq 2", "ab", true},
		{"len eq 3 not palindrome", "abc", false},
		{"len eq 3", "aab", true},
		{"abca", "abca", true},
		{"abcb", "abcb", true},
		{"eedede", "eedede", true},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := ValidPalindrome(c.s)
			if ret != c.exp {
				t.Fatalf("expeced %t but got %t, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
