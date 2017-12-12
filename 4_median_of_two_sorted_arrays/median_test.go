package median

import "testing"

func TestMedian(t *testing.T) {
	st := []struct {
		name         string
		nums1, nums2 []int
		exp          float64
	}{
		{"empty", []int{}, []int{}, 0.0},
		{"one is empty", []int{}, []int{1, 2, 3}, 2.0},
		{"test 1", []int{1, 3}, []int{2}, 2.0},
		{"test 2", []int{1, 2}, []int{3, 4}, 2.5},
		{"failed 1", []int{1, 1}, []int{1, 1}, 1.0},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Median(c.nums1, c.nums2)
			if ret != c.exp {
				t.Fatalf("expected %.2f but got %.2f with input %v and %v",
					c.exp, ret, c.nums1, c.nums2)
			}
		})
	}
}

func TestMedian2(t *testing.T) {
	st := []struct {
		name         string
		nums1, nums2 []int
		exp          float64
	}{
		{"empty", []int{}, []int{}, 0.0},
		{"one is empty", []int{}, []int{1, 2, 3}, 2.0},
		{"test 1", []int{1, 3}, []int{2}, 2.0},
		{"test 2", []int{1, 2}, []int{3, 4}, 2.5},
		{"failed 1", []int{1, 1}, []int{1, 1}, 1.0},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Median2(c.nums1, c.nums2)
			if ret != c.exp {
				t.Fatalf("expected %.2f but got %.2f with input %v and %v",
					c.exp, ret, c.nums1, c.nums2)
			}
		})
	}
}
