package fib

import "testing"

func TestFib(t *testing.T) {
	st := []struct {
		name string
		N    int
		exp  int
	}{
		{"n=0", 0, 0},
		{"n=1", 1, 1},
		{"n=2", 2, 1},
		{"n=4", 4, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := fib(tt.N)
			if out != tt.exp {
				t.Fatalf("with input n:%d ,wanted %d but got %d", tt.N, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
