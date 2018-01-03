package utils

import (
	"reflect"
	"testing"
)

func TestDeleteFromIntSlice(t *testing.T) {
	st := []struct {
		name   string
		a      []int
		target int
		exp    []int
	}{
		{"empty slice", []int{}, 1, []int{}},
		{"single", []int{1}, 2, []int{1}},
		{"test1", []int{1, 2, 3}, 2, []int{1, 3}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := DeleteFromIntSlice(c.a, c.target)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v and %d",
					c.exp, ret, c.a, c.target)
			}
		})
	}
}
