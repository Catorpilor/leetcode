package is

import "testing"

func TestIsomorphic(t *testing.T) {
	st := []struct {
		name string
		s    string
		t    string
		exp  bool
	}{
		{"egg and add", "egg", "add", true},
		{"foo and bar", "foo", "bar", false},
		{"title and paper", "paper", "title", true},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Isomorphic(c.s, c.t)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t with input %s and %s",
					c.exp, ret, c.s, c.t)
			}
		})
	}
}
