package gp

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  []string
	}{
		{"n=0", 0, []string{}},
		{"n=1", 1, []string{"()"}},
		{"n=2", 2, []string{"()()", "(())"}},
		{"n=3", 3, []string{"()()()", "((()))", "(())()", "()(())", "(()())"}},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := generateParenthesis(tt.n)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input n: %d wanted %v but got %v", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
