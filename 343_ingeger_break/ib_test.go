package ib

import "testing"

func TestIB(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name   string
		n, exp int
	}{
		{"n eq 2", 2, 1},
		{"n eq 4", 4, 4},
		{"n eq 6", 6, 9},
		{"n eq 7", 7, 12},
		{"n eq 9", 9, 27},
		{"n eq 11", 11, 54},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := IB(c.n)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}
func TestIB2(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name   string
		n, exp int
	}{
		{"n eq 2", 2, 1},
		{"n eq 4", 4, 4},
		{"n eq 6", 6, 9},
		{"n eq 7", 7, 12},
		{"n eq 9", 9, 27},
		{"n eq 11", 11, 54},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := IB2(c.n)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}
