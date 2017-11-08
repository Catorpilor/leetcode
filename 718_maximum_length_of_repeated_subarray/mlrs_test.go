package mlrs

import "testing"

func TestMaxLength(t *testing.T) {
	st := []struct {
		name         string
		nums1, nums2 []int
		exp          int
	}{
		{"nil nums1", nil, []int{1, 2, 3}, 0},
		{"nil nums2", []int{1, 2, 3}, nil, 0},
		{"single element", []int{1}, []int{1, 2, 3}, 1},
		{"test 1", []int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
		{"failed 1", []int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxLength(c.nums1, c.nums2)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d ,with input %v and %v",
					c.exp, ret, c.nums1, c.nums2)
			}
		})
	}
}
