package pow

import "testing"

func TestPow(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  bool
	}{
		{"n eq to 0", 0, false},
		{"n eq to 1", 1, true},
		{"n eq to 127", 127, false},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Pow3(c.n)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}
