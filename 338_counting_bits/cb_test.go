package cb

import (
	"reflect"
	"testing"
)

func TestCoutingBits(t *testing.T) {
	st := []struct {
		name string
		n    int
		exp  []int
	}{
		{"n=0", 0, []int{0}},
		{"n=1", 1, []int{0, 1}},
		{"n=3", 3, []int{0, 1, 1, 2}},
		{"n=5", 5, []int{0, 1, 1, 2, 1, 2}},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CountingBits(c.n)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with inputs %d",
					c.exp, ret, c.n)
			}
		})
	}
}
