package coins

import "testing"

func TestArrangeCoins(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"n=1", 1, 1},
		{"n=0", 0, 0},
		{"n=5", 5, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := arrangeCoins(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestArrangeCoinsUseBs(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"n=1", 1, 1},
		{"n=0", 0, 0},
		{"n=5", 5, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useBinarySearch(tt.n)
			if out != tt.exp {
				t.Fatalf("with input n:%d wanted %d but got %d", tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
