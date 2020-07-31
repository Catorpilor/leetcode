package climb

import "testing"

func TestNumOfClimbs(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"n=0", 0, 0},
		{"n=1", 1, 1},
		{"n=2", 2, 2},
		{"n=3", 3, 3},
		{"n=4", 4, 5},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := numOfClimbs(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d  but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
