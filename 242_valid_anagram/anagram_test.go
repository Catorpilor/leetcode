package anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	st := []struct {
		name string
		s, t string
		exp  bool
	}{
		{"empty string", "", "", true},
		{"normal case", "ab", "ba", true},
		{"not anagram", "abc", "aac", false},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := IsAnagram2(c.s, c.t)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %s and %s",
					c.exp, ret, c.s, c.t)
			}
		})
	}
}
