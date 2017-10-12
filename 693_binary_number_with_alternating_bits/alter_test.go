package alter

import "testing"

func TestHasAlter(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  bool
	}{
		{"5 has alter", 5, true},
		{"4 is not", 4, false},
		{"7 is not", 7, false},
		{"11 is not", 11, false},
		{"10 has alter", 10, true},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := HasAlter(c.n)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with inputs %d",
					c.exp, ret, c.n)
			}
		})
	}
}

func TestHasAlter2(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  bool
	}{
		{"5 has alter", 5, true},
		{"4 is not", 4, false},
		{"7 is not", 7, false},
		{"11 is not", 11, false},
		{"10 has alter", 10, true},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := HasAlter2(c.n)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with inputs %d",
					c.exp, ret, c.n)
			}
		})
	}
}
