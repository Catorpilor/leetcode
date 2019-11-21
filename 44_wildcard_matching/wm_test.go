package wm

import "testing"

func TestIsMatchTLE(t *testing.T) {
	st := []struct {
		name string
		s, p string
		exp  bool
	}{
		{"empty s", "", "b", false},
		{"empty p", "a", "", false},
		{"both empty", "", "", true},
		{"just alphabets", "abc", "ab", false},
		{"with question", "abc", "ab?", true},
		{"with asterisk", "abc", "*", true},
		{"all asterisk longer", "abc", "*************", true},
		{"testcase1", "adceb", "*a*b", true},
		{"failed1", "ba", "*a*", true},
		{"failed2", "hello", "*o*", true},
		{"failed3", "mississippi", "m*si*", true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isMatchTLE(tt.s, tt.p)
			if out != tt.exp {
				t.Fatalf("with s: %s and p: %s wanted %t but got %t", tt.s, tt.p, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestIsMatch(t *testing.T) {
	st := []struct {
		name string
		s, p string
		exp  bool
	}{
		{"empty s", "", "b", false},
		{"empty p", "a", "", false},
		{"both empty", "", "", true},
		{"just alphabets", "abc", "ab", false},
		{"with question", "abc", "ab?", true},
		{"with asterisk", "abc", "*", true},
		{"all asterisk longer", "abc", "*************", true},
		{"testcase1", "adceb", "*a*b", true},
		{"failed1", "ba", "*a*", true},
		{"failed2", "hello", "*o*", true},
		{"failed3", "mississippi", "m*si*", true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isMatch(tt.s, tt.p)
			if out != tt.exp {
				t.Fatalf("with s: %s and p: %s wanted %t but got %t", tt.s, tt.p, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
