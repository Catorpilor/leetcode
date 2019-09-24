package nge

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	st := []struct {
		name         string
		nums1, nums2 []int
		exp          []int
	}{
		{"both empty", []int{}, []int{}, []int{}},
		{"nums1 empty", []int{}, []int{1, 3, 4}, []int{}},
		{"nums2 empty", []int{1, 3}, []int{}, []int{-1, -1}},
		{"test1", []int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{"test2", []int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := nextGreaterElement(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input nums1: %v and nums2: %v, wanted %v but got %v", tt.nums1, tt.nums2, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
