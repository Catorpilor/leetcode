package set

import (
	"reflect"
	"testing"
)

func TestFindNums(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"two 1s", []int{1, 1}, []int{1, 2}},
		{"two 3s", []int{1, 3, 3}, []int{3, 2}},
		{"two 2s lenght 4", []int{4, 2, 1, 2}, []int{2, 3}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := FindNums(c.nums)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestFindNums2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"two 1s", []int{1, 1}, []int{1, 2}},
		{"two 3s", []int{1, 3, 3}, []int{3, 2}},
		{"two 2s lenght 4", []int{4, 2, 1, 2}, []int{2, 3}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := FindNums2(c.nums)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
