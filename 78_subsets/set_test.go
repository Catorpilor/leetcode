package set

import (
	"reflect"
	"testing"
)

func TestSubSets(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  [][]int
	}{
		{"single element", []int{1}, [][]int{[]int{}, []int{1}}},
		{"two elements", []int{1, 2}, [][]int{[]int{}, []int{1}, []int{2}, []int{1, 2}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := subSets(c.nums)
			if !reflect.DeepEqual(c.exp, ret) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}

}

func TestSubSets2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  [][]int
	}{
		{"single element", []int{1}, [][]int{[]int{}, []int{1}}},
		{"two elements", []int{1, 2}, [][]int{[]int{}, []int{1}, []int{2}, []int{1, 2}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := SubSets2(c.nums)
			if len(ret) != len(c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}

}
