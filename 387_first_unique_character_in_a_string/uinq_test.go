package uniq

import "testing"

func TestFirstUniq(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  int
	}{
		{"empty string", "", -1},
		{"one character", "a", 0},
		{"double charaters without duplicates", "ab", 0},
		{"double charaters with duplicates", "aa", -1},
		{"leetcode", "leetcode", 0},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := FirstUniq(c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
