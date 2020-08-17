package eat

import "testing"

func TestMinDays(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"n=1", 1, 1},
		{"n=2", 2, 2},
		{"n=3", 3, 2},
		{"n=5", 5, 4},
		{"n=10", 10, 4},
		{"n=56", 56, 6},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minDays(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
