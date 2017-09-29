package palin

import "testing"

func TestCanPermutePalindrome(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  bool
	}{
		{"empty", "", true},
		{"single character", "a", true},
		{"len(s) eq 2", "ab", false},
		{"normal", "code", false},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CanPermutePalindrome2(c.s)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
