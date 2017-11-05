package seq

import "testing"

func TestIsSub(t *testing.T) {
	st := []struct {
		name string
		s, t string
		exp  bool
	}{
		{"empty s", "", "abc", true},
		{"empty s and t", "", "", true},
		{"empty t, non-empty s", "a", "", false},
		{"s longer than t", "abc", "a", false},
		{"abc, ahbgdc", "abc", "ahbgcd", true},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := IsSub(c.s, c.t)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %s and %s",
					c.exp, ret, c.s, c.t)
			}
		})
	}
}

func TestIsSubBS(t *testing.T) {
	st := []struct {
		name string
		s, t string
		exp  bool
	}{
		{"empty s", "", "abc", true},
		{"empty s and t", "", "", true},
		{"empty t, non-empty s", "a", "", false},
		{"s longer than t", "abc", "a", false},
		{"abc ahbgdc", "abc", "ahbgcd", true},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := IsSubBS(c.s, c.t)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %s and %s",
					c.exp, ret, c.s, c.t)
			}
		})
	}
}
