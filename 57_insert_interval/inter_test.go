package inter

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	st := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		exp         [][]int
	}{
		{"empty slice", nil, []int{1, 2}, [][]int{[]int{1, 2}}},
		{"single", [][]int{[]int{1, 3}}, []int{2, 4}, [][]int{[]int{1, 4}}},
		{"single", [][]int{[]int{1, 5}}, []int{2, 4}, [][]int{[]int{1, 5}}},
		{"test 1", [][]int{[]int{1, 3}, []int{6, 9}}, []int{2, 5}, [][]int{[]int{1, 5}, []int{6, 9}}},
		{"test 2", [][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}},
			[]int{4, 9}, [][]int{[]int{1, 2}, []int{3, 10}, []int{12, 16}}},
		{"failed 1", [][]int{[]int{1, 5}}, []int{0, 3}, [][]int{[]int{0, 5}}},
		{"failed 2", [][]int{[]int{1, 5}}, []int{0, 6}, [][]int{[]int{0, 6}}},
		{"failed 3", [][]int{[]int{1, 5}, []int{6, 8}}, []int{0, 9}, [][]int{[]int{0, 9}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := insertIntervals(c.intervals, c.newInterval)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v and %v",
					c.exp, ret, c.intervals, c.newInterval)
			}
		})
	}
}
