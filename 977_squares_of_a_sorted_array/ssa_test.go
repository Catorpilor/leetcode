package ssa

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	st := []struct {
		name string
		A    []int
		exp  []int
	}{
		{"empty array", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"identical", []int{-1, -1, -1}, []int{1, 1, 1}},
		{"test1", []int{-4, -2, 0, 1, 3, 5}, []int{0, 1, 4, 9, 16, 25}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := sortedSquares(tt.A)
			if !reflect.DeepEqual(out, tt.exp) {
				t.Fatalf("with ipnut A:%v wanted %v but got %v", tt.A, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
