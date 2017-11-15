package wrap

import "testing"

func TestNumOfSubStrings(t *testing.T) {
	st := []struct {
		name, p string
		exp     int
	}{
		{"empty string", "", 0},
		{"identical", "aa", 1},
		{"single", "a", 1},
		{"cac", "cac", 2},
		{"zab", "zab", 6},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := NumOfSubStrings(c.p)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %s",
					c.exp, ret, c.p)
			}
		})
	}
}
func TestNumOfSubStrings2(t *testing.T) {
	st := []struct {
		name, p string
		exp     int
	}{
		{"empty string", "", 0},
		{"identical", "aa", 1},
		{"single", "a", 1},
		{"cac", "cac", 2},
		{"zab", "zab", 6},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := NumOfSubStrings2(c.p)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %s",
					c.exp, ret, c.p)
			}
		})
	}
}
