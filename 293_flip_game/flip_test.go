package flip

import (
	"reflect"
	"testing"
)

func TestFlip(t *testing.T) {
	st := []struct {
		name, s string
		exp     []string
	}{
		{"++++", "++++", []string{"--++", "+--+", "++--"}},
		{"++--", "++--", []string{"----"}},
		{"+-+-", "+-+-", []string{}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Flip(c.s)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
