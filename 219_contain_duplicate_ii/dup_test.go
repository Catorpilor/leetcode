package dup

import "testing"

func TestDup(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  bool
	}{
		{"nil slice", nil, 1, false},
		{"single slice", []int{1}, 1, false},
		{"failed once", []int{1, 0, 1, 1}, 1, true},
		{"normal case without duplicate", []int{1, 2, 3}, 1, false},
		{"nc with duplicate", []int{1, 2, 2}, 1, true},
		{"nc with duplicate not in k distance", []int{2, 1, 2}, 1, false},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Dup(c.nums, c.k)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with inputs %v and %d",
					c.exp, ret, c.nums, c.k)
			}
		})
	}
}
