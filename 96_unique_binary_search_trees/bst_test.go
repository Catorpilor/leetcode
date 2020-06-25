package bst

import "testing"

func TestNumberOfBst(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"n=1", 1, 1},
		{"n=2", 2, 2},
		{"n=3", 3, 5},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := numOfBst(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
