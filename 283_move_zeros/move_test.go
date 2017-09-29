package move

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"nil slice", nil, nil},
		{"all zeroes", []int{0, 0, 0, 0}, []int{0, 0, 0, 0}},
		{"normal case", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"last idx is not zero", []int{0, 0, 0, 1}, []int{1, 0, 0, 0}},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MoveZeroes(c.nums)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but go %v, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
