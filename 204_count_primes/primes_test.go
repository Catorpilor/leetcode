package primes

import "testing"

func TestCountPrimes(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  int
	}{
		{"less than 2", 2, 0},
		{"n eq 5", 5, 2},
		{"n eq 6", 6, 3},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CountPrimes2(c.n)
			if ret != c.exp {
				t.Fatalf("expeced %d but got %d, with input %d",
					c.exp, ret, c.n)
			}
		})
	}
}
