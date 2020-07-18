package stone

import "testing"

func TestWinnerSquareGame(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  bool
	}{
		{"n=1", 1, true},
		{"n=2", 2, false},
		{"n=3", 3, true},
		{"n=4", 4, true},
		{"n=5", 5, false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := winnerSquareGame(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %t but got %t", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
