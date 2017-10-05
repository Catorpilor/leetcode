package num

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"missing 5,6", []int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
		{"missing 3", []int{1, 1, 2}, []int{3}},
		{"no missing", []int{1, 2, 3}, nil},
		{"extrame", []int{1, 1, 1, 1, 1, 1}, []int{2, 3, 4, 5, 6}},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := FindDisappearedNumbers(c.nums)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
