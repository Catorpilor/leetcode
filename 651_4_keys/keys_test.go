package keys

import "testing"

func TestMaxA(t *testing.T) {
	st := []struct {
		name   string
		n, exp int
	}{
		{"n eq 0", 0, 0},
		{"n eq 1", 1, 1},
		{"n eq 2", 2, 2},
		{"n eq 6", 6, 6},
		{"n eq 7", 7, 9},
		{"n eq 8", 8, 12},
		{"n eq 10", 10, 20},
		{"n eq 12", 12, 36},
		{"n eq 13", 13, 48},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxA(c.n)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}

func TestMaxA2(t *testing.T) {
	st := []struct {
		name   string
		n, exp int
	}{
		{"n eq 0", 0, 0},
		{"n eq 1", 1, 1},
		{"n eq 2", 2, 2},
		{"n eq 6", 6, 6},
		{"n eq 7", 7, 9},
		{"n eq 8", 8, 12},
		{"n eq 10", 10, 20},
		{"n eq 12", 12, 36},
		{"n eq 13", 13, 48},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxA2(c.n)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}

func TestMaxA3(t *testing.T) {
	st := []struct {
		name   string
		n, exp int
	}{
		{"n eq 0", 0, 0},
		{"n eq 1", 1, 1},
		{"n eq 2", 2, 2},
		{"n eq 6", 6, 6},
		{"n eq 7", 7, 9},
		{"n eq 8", 8, 12},
		{"n eq 10", 10, 20},
		{"n eq 12", 12, 36},
		{"n eq 13", 13, 48},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxA3(c.n)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}
