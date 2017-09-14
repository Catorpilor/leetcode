package dif

import "testing"

func TestDif(t *testing.T) {
	st := []struct {
		name string
		s    string
		t    string
		exp  byte
	}{
		{"empty s", "", "t", byte('t')},
		{"single s", "a", "aa", byte('a')},
		{"normal s", "abcd", "dbcae", byte('e')},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Dif2(c.s, c.t)
			if ret != c.exp {
				t.Fatalf("expected %s but got %s ,with inputs %s and %s",
					string(c.exp), string(ret), c.s, c.t)
			}
		})
	}
}
