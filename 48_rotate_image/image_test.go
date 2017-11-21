package image

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	st := []struct {
		name   string
		matrix [][]int
		exp    [][]int
	}{
		{"empty matrix", [][]int{}, [][]int{}},
		{"1*1", [][]int{[]int{1}}, [][]int{[]int{1}}},
		{"2*2", [][]int{[]int{1, 2}, []int{3, 4}}, [][]int{[]int{3, 1}, []int{4, 2}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Rotate(c.matrix)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %v",
					c.exp, ret, c.matrix)
			}
		})
	}
}
