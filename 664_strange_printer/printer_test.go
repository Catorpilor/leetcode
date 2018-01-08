package printer

import "testing"

func TestStrangePrinter(t *testing.T) {
	st := []struct {
		name, s string
		exp     int
	}{
		{"empty", "", 0},
		{"single", "a", 1},
		{"identicals", "aaaaa", 1},
		{"test1", "aaaabbb", 2},
		{"test2", "aba", 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := StrangePrinter(c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}

func TestStrangePrinter2(t *testing.T) {
	st := []struct {
		name, s string
		exp     int
	}{
		{"empty", "", 0},
		{"single", "a", 1},
		{"identicals", "aaaaa", 1},
		{"test1", "aaaabbb", 2},
		{"test2", "aba", 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := StrangePrinter2(c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}

func TestStrangePrinter3(t *testing.T) {
	st := []struct {
		name, s string
		exp     int
	}{
		{"empty", "", 0},
		{"single", "a", 1},
		{"identicals", "aaaaa", 1},
		{"test1", "aaaabbb", 2},
		{"test2", "aba", 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := StrangePrinter3(c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
