package permuate

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	st := []struct {
		name      string
		nums, exp []int
	}{
		{"nil slice", nil, nil},
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"test 123", []int{1, 2, 3}, []int{1, 3, 2}},
		{"test 1231", []int{1, 2, 3, 1}, []int{1, 3, 1, 2}},
		{"test 1312", []int{1, 3, 1, 2}, []int{1, 3, 2, 1}},
		{"test 1321", []int{1, 3, 2, 1}, []int{2, 1, 1, 3}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := NextPermutation(c.nums)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}

}
