package dev

import "testing"

func TestMinimumDeviation(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"testcase1", []int{1, 2, 3, 4}, 1},
		{"testcase2", []int{4, 1, 5, 20, 3}, 3},
		{"testcase3", []int{2, 10, 8}, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minimumDeviation(tt.nums)
			if out != tt.exp {
				t.Fatalf("wanted %d but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
