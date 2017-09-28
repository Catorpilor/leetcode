package add

import "testing"

func TestAddDigits(t *testing.T) {
	st := []struct {
		name string
		num  int
		exp  int
	}{
		{"num < 10", 8, 8},
		{"38", 38, 2},
		{"18", 18, 9},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := AddDigits(c.num)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %d",
					c.exp, ret, c.num)
			}
		})
	}
}
