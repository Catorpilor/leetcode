package product

import (
	"reflect"
	"testing"
)

func TestProductOfArray(t *testing.T) {
	st := []struct {
		name      string
		nums, exp []int
	}{
		{"nil slice", nil, nil},
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"identical two", []int{1, 1}, []int{1, 1}},
		{"test 1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := ProductOfArray(c.nums)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestProductOfArray2(t *testing.T) {
	st := []struct {
		name      string
		nums, exp []int
	}{
		{"nil slice", nil, nil},
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"identical two", []int{1, 1}, []int{1, 1}},
		{"test 1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := ProductOfArray2(c.nums)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
