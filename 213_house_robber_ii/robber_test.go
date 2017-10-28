package robber

import "testing"

func TestRobber(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 1},
		{"two elements", []int{1, 2}, 2},
		{"3 elements", []int{1, 2, 3}, 3},
		{"test 31524", []int{3, 1, 5, 2, 4}, 9},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Robber(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}

}
