package lp

import "testing"

func TestUseBucket(t *testing.T) {
	st := []struct {
		name, s string
		exp     int
	}{
		{"empty string", "", 0},
		{"abcccdd", "abcccdd", 5},
		{"ccc", "ccc", 3},
		{"abccccdd", "abccccdd", 7},
		{"AAAAAA", "AAAAAA", 6},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useBucket(c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
