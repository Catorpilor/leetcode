package combine

import "testing"

func TestCombination(t *testing.T) {
	st := []struct {
		name        string
		nums        []int
		target, exp int
	}{
		{"single element 1", []int{1}, 4, 1},
		{"single element 2", []int{2}, 3, 0},
		{"test 123", []int{1, 2, 3}, 4, 7},
		{"test 213", []int{2, 1, 3}, 3, 4},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Combination(c.nums, c.target)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v and %d",
					c.exp, ret, c.nums, c.target)
			}
		})
	}
}
