package remove

import (
	"testing"
)

func TestRemoveInvalidP(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  []string
	}{
		{"empty", "", []string{}},
		{"single", "(", []string{}},
		{"test 1", "()())()", []string{"()()()", "(())()"}},
		{"test2", "(a)())()", []string{"(a)()()", "(a())()"}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := RemoveInvalidP(c.s)
			if len(ret) != len(c.exp) {
				t.Fatalf("expected %v but got %v with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
