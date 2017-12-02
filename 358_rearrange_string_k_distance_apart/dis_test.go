package dis

import "testing"

func TestRearrange(t *testing.T) {
	st := []struct {
		name, s string
		k       int
		exp     string
	}{
		{"empty string", "", 2, ""},
		{"single str", "a", 3, "a"},
		{"two identical", "aa", 2, ""},
		{"test 1", "aaadbbcc", 2, "abacabcd"},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Rearrange(c.s, c.k)
			if ret != c.exp {
				t.Fatalf("expected %s but got %s with input %s and %d",
					c.exp, ret, c.s, c.k)
			}
		})
	}
}
