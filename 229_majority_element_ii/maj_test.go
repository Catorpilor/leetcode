package major

import (
	"reflect"
	"testing"
)

func TestMajority(t *testing.T) {
	st := []struct {
		name      string
		nums, exp []int
	}{
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{1, 2}},
		{"three elements", []int{1, 2, 3}, []int{}},
		{"four elements", []int{1, 2, 2, 3}, []int{2}},
		{"five elements", []int{1, 2, 2, 1, 3}, []int{1, 2}},
		{"failed 1", []int{2, 2}, []int{2}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Majority(c.nums)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestMajority2(t *testing.T) {
	st := []struct {
		name      string
		nums, exp []int
	}{
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"three elements", []int{1, 2, 3}, []int{}},
		{"four elements", []int{1, 2, 2, 3}, []int{2}},
		{"five elements", []int{1, 2, 2, 1, 3}, []int{2, 1}},
		{"failed 1", []int{2, 2}, []int{2}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Majority2(c.nums)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
