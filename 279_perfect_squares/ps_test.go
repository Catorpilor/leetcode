package ps

import "testing"

func TestNumSquares(t *testing.T) {
	st := []struct {
		name string
		num  int
		exp  int
	}{
		{"num eq 0", 0, 0},
		{"num eq 1", 1, 1},
		{"num eq 2", 2, 2},
		{"num eq 4", 5, 2},
		{"num eq 9", 9, 1},
		{"num eq 12", 12, 3},
		{"num eq 13", 13, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := NumSquares(c.num)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with inputs %d",
					c.exp, ret, c.num)
			}
		})
	}
}
