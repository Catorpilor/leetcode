package permutate

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  [][]int
	}{
		{"empty slice", []int{}, [][]int{[]int{}}},
		{"single element", []int{1}, [][]int{[]int{1}}},
		{"two elements", []int{1, 2}, [][]int{[]int{1, 2}, []int{2, 1}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Permute(c.nums)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestPermute2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  [][]int
	}{
		{"empty slice", []int{}, [][]int{[]int{}}},
		{"single element", []int{1}, [][]int{[]int{1}}},
		{"two elements", []int{1, 2}, [][]int{[]int{1, 2}, []int{2, 1}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Permute2(c.nums)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestPermuteUsebacktrack(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  [][]int
	}{
		{"empty slice", []int{}, [][]int{[]int{}}},
		{"single element", []int{1}, [][]int{[]int{1}}},
		{"two elements", []int{1, 2}, [][]int{[]int{1, 2}, []int{2, 1}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := permute(c.nums)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
