package merge

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	st := []struct {
		name           string
		intervals, exp []interval
	}{
		{"empty intervals", []interval{}, []interval{}},
		{"no overlapping", []interval{interval{1, 2}, interval{3, 4}},
			[]interval{interval{1, 2}, interval{3, 4}}},
		{"with overlapping", []interval{interval{1, 3}, interval{2, 6}}, []interval{interval{1, 6}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Merge(c.intervals)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %v",
					c.exp, ret, c.intervals)
			}
		})
	}
}
