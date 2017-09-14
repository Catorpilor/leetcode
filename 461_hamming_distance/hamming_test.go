package hamming

import "testing"

func TestDistance(t *testing.T) {
	st := []struct {
		name string
		x    int
		y    int
		exp  int
	}{
		{"same numbers", 1, 1, 0},
		{"1 and 4", 1, 4, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Distance(c.x, c.y)
			if ret != c.exp {
				t.Fatalf("expected %d bug got %d, with input %d and %d",
					c.exp, ret, c.x, c.y)
			}
		})
	}
}
