package score

import "testing"

func TestScoreOfParenthese(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  int
	}{
		{"invalid pair", ")(", 0},
		{"testcase1", "()", 1},
		{"testcase2", "()()", 2},
		{"testcase3", "(())", 2},
		{"testcase4", "(()(()))", 6},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := scoreOfParenthese(tt.s)
			if out != tt.exp {
				t.Fatalf("with input s:%s wanted %d but got %d", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
