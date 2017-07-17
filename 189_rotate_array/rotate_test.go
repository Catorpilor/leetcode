package rotate

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		k      int
		expect []int
	}{
		{"nil slice", nil, -1, nil},
		{"empty slice", []int{}, 1, []int{}},
		{"1 element", []int{1}, 1, []int{1}},
		{"k is bigger than the length", []int{1, 2}, 3, []int{2, 1}},
		{"normal", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Rotate(c.nums, c.k)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v, but got %v, with inputs %v and k %d",
					c.expect, ret, c.nums, c.k)
			}
		})
	}
	t.Log("passed")
}
