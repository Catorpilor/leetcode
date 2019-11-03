package unique

import "testing"

func TestCountNumbers(t *testing.T) {
	st := []struct {
		name   string
		n, exp int
	}{
		{"n=0", 0, 1},
		{"n=1", 1, 10},
		{"n=2", 2, 91},
		{"n=3", 3, 739},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := countUniqueNumbers(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestBackTracking(t *testing.T) {
	st := []struct {
		name   string
		n, exp int
	}{
		{"n=0", 0, 1},
		{"n=1", 1, 10},
		{"n=2", 2, 91},
		{"n=3", 3, 739},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := backtracking(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
