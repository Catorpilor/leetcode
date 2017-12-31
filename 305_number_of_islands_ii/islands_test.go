package islands

import (
	"reflect"
	"testing"
)

func TestNumberOfIslands(t *testing.T) {
	st := []struct {
		name string
		m, n int
		pos  [][]int
		exp  []int
	}{
		{"0*0", 0, 0, nil, nil},
		{"1*0", 1, 0, nil, nil},
		{"pos is nil", 3, 3, nil, nil},
		{"test1", 3, 3, [][]int{[]int{0, 0}, []int{0, 1}, []int{1, 2}, []int{2, 1}}, []int{1, 1, 2, 3}},
		{"failed1", 2, 2, [][]int{[]int{0, 0}, []int{1, 1}, []int{0, 1}}, []int{1, 2, 1}},
		{"failed2", 3, 3, [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 1}, []int{1, 0}, []int{0, 2}, []int{0, 0}, []int{1, 1}}, []int{1, 2, 3, 4, 3, 2, 1}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := NumberOfIslands(c.m, c.n, c.pos)
			if !reflect.DeepEqual(c.exp, ret) {
				t.Fatalf("expected %v but got %v with input %d and %d and %v",
					c.exp, ret, c.m, c.n, c.pos)
			}
		})
	}
}
