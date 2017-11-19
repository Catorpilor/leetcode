package search

import "testing"

func TestIsExist(t *testing.T) {
	st := []struct {
		name        string
		nums        []int
		target, exp int
	}{
		{"empty slice", []int{}, 5, -1},
		{"single slice", []int{1}, 2, -1},
		{"test 1", []int{4, 5, 6, 7, 0, 1, 2, 3}, 7, 3},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := IsExist(c.nums, c.target)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v and %d",
					c.exp, ret, c.nums, c.target)
			}
		})
	}
}
