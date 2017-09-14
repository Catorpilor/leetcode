package happy

import "testing"

func TestIsHappy(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  bool
	}{
		{"zero", 0, false},
		{"one", 1, true},
		{"19", 19, true},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := IsHappy(c.n)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t ,with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}
