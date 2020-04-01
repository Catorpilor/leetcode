package pn

import "testing"

func TestCountSubStrings(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name, s string
		exp     int
	}{
		{"empty string", "", 0},
		{"single character", "a", 1},
		{"two characters not eq", "ab", 2},
		{"two identifcal chars", "aa", 3},
		{"test ababb", "ababb", 8},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CountSubStrings(c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}

func TestUseBruteForce(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name, s string
		exp     int
	}{
		{"empty string", "", 0},
		{"single character", "a", 1},
		{"two characters not eq", "ab", 2},
		{"two identifcal chars", "aa", 3},
		{"test ababb", "ababb", 8},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useBruteForce(c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}