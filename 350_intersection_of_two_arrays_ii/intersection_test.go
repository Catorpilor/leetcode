package intersection

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	st := []struct {
		name   string
		nums1  []int
		nums2  []int
		expect []int
	}{
		{"all nil", nil, nil, nil},
		{"one is nil", nil, []int{1}, nil},
		{"non nil without intersection", []int{1, 2}, []int{3, 4}, nil},
		{"non nil with intersection", []int{1, 2, 3}, []int{3, 4}, []int{3}},
		{"duplicates", []int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Intersection(c.nums1, c.nums2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v, but got %v, with inputs %v and %v",
					c.expect, ret, c.nums1, c.nums2)
			}
		})
	}
}
