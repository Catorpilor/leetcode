package sum

import (
	"reflect"
	"testing"
)

func TestMaxSumOfThree(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  []int
	}{
		{"empty slice", []int{}, 2, []int{}},
		{"n/3 is less than 1", []int{1, 2}, 1, []int{}},
		{"k is greater than n/3", []int{1, 2, 3, 1, 2}, 3, []int{}},
		{"test 1", []int{1, 2, 1, 2, 6, 7, 5, 1}, 2, []int{0, 3, 5}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxSumOfThree(c.nums, c.k)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v and %d", c.exp, ret, c.nums, c.k)
			}
		})
	}
}
