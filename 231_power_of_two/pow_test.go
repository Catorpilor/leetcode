package pow

import "testing"

func TestPow(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  bool
	}{
		{"n eq 0", 0, false},
		{"n eq 1", 1, true},
		{"n eq 2", 2, true},
		{"n eq 3", 3, false},
		{"n eq 144", 144, false},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Pow(c.n)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}
