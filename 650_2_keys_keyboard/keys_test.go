package keys

import "testing"

func TestMinStep(t *testing.T) {
	st := []struct {
		name   string
		n, exp int
	}{
		{"n=0", 0, 0},
		{"n=1", 1, 0},
		{"n=2", 2, 2},
		{"n=5", 5, 5},
		{"n=6", 6, 5},
		{"n=18", 18, 8},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MinStep(c.n)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}

func TestMinStep2(t *testing.T) {
	st := []struct {
		name   string
		n, exp int
	}{
		{"n=0", 0, 0},
		{"n=1", 1, 0},
		{"n=2", 2, 2},
		{"n=5", 5, 5},
		{"n=6", 6, 5},
		{"n=18", 18, 8},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MinStep2(c.n)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}
