package nqueue

import "testing"

func TestTotalNQueues(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"n=1", 1, 1},
		{"n=2", 2, 0},
		{"n=3", 3, 0},
		{"n=4", 4, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := totalNQueues(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n: %d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
