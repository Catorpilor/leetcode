package dup

import "testing"

func TestDup(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  bool
	}{
		{"nil slice", nil, false},
		{"single slice", []int{1}, false},
		{"normal case without dup", []int{1, 2, 3}, false},
		{"normal case with dups", []int{1, 2, 2}, true},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Dup(c.nums)
			if ret != c.exp {
				t.Fatalf("expeced %t but got %t ,with inputs %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
