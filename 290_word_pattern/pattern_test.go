package pattern

import "testing"

func TestWordPattern(t *testing.T) {
	st := []struct {
		name    string
		pattern string
		str     string
		exp     bool
	}{
		{"abba", "abba", "dog cat cat dog", true},
		{"abba not match", "abba", "dog cat cat fish", false},
		{"aaaa not match", "aaaa", "dog cat cat cat", false},
		{"aaaa", "aaaa", "dog dog dog dog", true},
		{"aba not match", "aba", "dog cat cat", false},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := WordPattern(c.pattern, c.str)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %s and %s",
					c.exp, ret, c.pattern, c.str)
			}
		})
	}
}
