package up

import "testing"

func TestUniquePath(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name      string
		m, n, exp int
	}{
		{"m&n = 0", 0, 1, 0},
		{"m,n = 1", 3, 1, 1},
		{"n,m=1", 1, 3, 1},
		{"m=3,n=7", 3, 7, 28},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := UniquePath(c.m, c.n)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %d and %d",
					c.exp, ret, c.m, c.n)
			}
		})
	}
}
