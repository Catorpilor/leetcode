package reverse

import "testing"

func TestReverse(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  string
	}{
		{"emptry string", "", ""},
		{"single character", "a", "a"},
		{"normal case", "ab", "ba"},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Reverse(c.s)
			if ret != c.exp {
				t.Fatalf("expeced %s but got %s, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
