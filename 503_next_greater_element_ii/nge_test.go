package nge

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{-1}},
		{"identical", []int{1, 1, 1}, []int{-1, -1, -1}},
		{"decreasing", []int{5, 3, 2}, []int{-1, 5, 5}},
		{"increasing", []int{1, 2, 3}, []int{2, 3, -1}},
		{"testcase1", []int{1, 2, 1}, []int{2, -1, 2}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := nextGreaterElement(tt.nums)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input nums:%v wanted %v but got %v", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
