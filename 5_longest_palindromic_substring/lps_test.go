package lps

import "testing"

func TestLps(t *testing.T) {
	st := []struct {
		name, s, exp string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"two characters", "ab", "a"},
		{"two identical", "aa", "aa"},
		{"test ababb", "ababb", "bab"},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Lps(c.s)
			if ret != c.exp {
				t.Fatalf("expected %s but got %s, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
