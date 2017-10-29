package lps

import "testing"

func TestLps(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name, s string
		exp     int
	}{
		{"empty string", "", 0},
		{"single", "a", 1},
		{"two different chars", "ab", 1},
		{"two identical", "aa", 2},
		{"test bbbab", "bbbab", 4},
		{"test cbbd", "cbbd", 2},
		{"test aabaa", "aabaa", 5},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Lps(c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
