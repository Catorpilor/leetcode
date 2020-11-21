package nlis

import "testing"

func TestNumberOfLIS(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 1},
		{"two incresing", []int{1, 2}, 1},
		{"two decresing", []int{2, 1}, 2},
		{"identical", []int{2, 2, 2, 2, 2}, 5},
		{"test 1684", []int{1, 6, 8, 4}, 1},
		{"test 16847", []int{1, 6, 8, 4, 7}, 3},
		{"test 31547", []int{3, 1, 5, 4, 7}, 4},
		{"failed 10", []int{1, 2, 3, 1, 2, 3, 1, 2, 3}, 10},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := findNumberOfLIS(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
