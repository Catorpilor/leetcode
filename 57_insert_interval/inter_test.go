package inter

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	st := []struct {
		name        string
		intervals   []Interval
		newInterval Interval
		exp         []Interval
	}{
		{"empty slice", []Interval{}, Interval{1, 2}, []Interval{Interval{1, 2}}},
		{"single", []Interval{Interval{1, 3}}, Interval{2, 4}, []Interval{Interval{1, 4}}},
		{"single", []Interval{Interval{1, 5}}, Interval{2, 4}, []Interval{Interval{1, 5}}},
		{"test 1", []Interval{Interval{1, 3}, Interval{6, 9}}, Interval{2, 5}, []Interval{Interval{1, 5}, Interval{6, 9}}},
		{"test 2", []Interval{Interval{1, 2}, Interval{3, 5}, Interval{6, 7}, Interval{8, 10}, Interval{12, 16}},
			Interval{4, 9}, []Interval{Interval{1, 2}, Interval{3, 10}, Interval{12, 16}}},
		{"failed 1", []Interval{Interval{1, 5}}, Interval{0, 3}, []Interval{Interval{0, 5}}},
		{"failed 2", []Interval{Interval{1, 5}}, Interval{0, 6}, []Interval{Interval{0, 6}}},
		{"failed 3", []Interval{Interval{1, 5}, Interval{6, 8}}, Interval{0, 9}, []Interval{Interval{0, 9}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Insert(c.intervals, c.newInterval)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v and %v",
					c.exp, ret, c.intervals, c.newInterval)
			}
		})
	}
}
