package cc

import "testing"

func TestMaxPower(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  int
	}{
		{"empty string", "", 0},
		{"single charcter", "a", 1},
		{"testcase1", "leetcode", 2},
		{"testcase2", "abbcccddddeeeeedcba", 5},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxPower(tt.s)
			if out != tt.exp {
				t.Fatalf("with input s:%s wanted %d but got %d", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
