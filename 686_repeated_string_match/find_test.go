package find

import "testing"

func TestMatch3(t *testing.T) {
	st := []struct {
		name, a, b string
		exp        int
	}{
		{"both empty", "", "", -1},
		{"abcd", "abcd", "cdabcdab", 3},
		{"not found", "abcd", "edabcdab", -1},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Match3(c.a, c.b)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %s and %s",
					c.exp, ret, c.a, c.b)
			}
		})
	}
}
