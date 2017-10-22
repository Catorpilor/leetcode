package decode

import "testing"

func TestNumDecodings(t *testing.T) {
	st := []struct {
		name, s string
		exp     int
	}{
		{"empty string", "", 0},
		{"faield 0", "0", 0},
		{"faield 10", "10", 1},
		{"faield 100", "100", 0},
		{"faield 110", "110", 1},
		{"faield 1029", "1029", 1},
		{"single character", "1", 1},
		{"two characters", "12", 2},
		{"two characters", "27", 1},
		{"three characters", "123", 3},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := NumEncodings(c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}

func TestNumDecodings2(t *testing.T) {
	st := []struct {
		name, s string
		exp     int
	}{
		{"empty string", "", 0},
		{"faield 0", "0", 0},
		{"faield 10", "10", 1},
		{"faield 100", "100", 0},
		{"faield 110", "110", 1},
		{"faield 1029", "1029", 1},
		{"single character", "1", 1},
		{"two characters", "12", 2},
		{"two characters", "27", 1},
		{"three characters", "123", 3},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := NumEncodings2(c.s)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
