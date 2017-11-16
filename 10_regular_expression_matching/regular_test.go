package regular

import "testing"

func TestIsMatch(t *testing.T) {
	st := []struct {
		name, s, p string
		exp        bool
	}{
		{"empty s and p", "", "", true},
		{"one is empty", "a", "", false},
		{".* rules all", "abab", ".*", true},
		{"c.* fails", "aba", "c.*", false},
		{"aab", "aab", "c*a*b", true},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := IsMatch(c.s, c.p)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t ,with input %s and %s",
					c.exp, ret, c.s, c.p)
			}
		})
	}
}

func TestIsMatchDP(t *testing.T) {
	st := []struct {
		name, s, p string
		exp        bool
	}{
		{"empty s and p", "", "", true},
		{"one is empty", "a", "", false},
		{".* rules all", "abab", ".*", true},
		{"c.* fails", "aba", "c.*", false},
		{"aab", "aab", "c*a*b", true},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := IsMatchDP(c.s, c.p)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t ,with input %s and %s",
					c.exp, ret, c.s, c.p)
			}
		})
	}
}
