package array

import "testing"

func TestMinOperations(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"n=1", 1, 0},
		{"n=2", 2, 1},
		{"n=3", 3, 2},
		{"n=6", 6, 9},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minOps(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
