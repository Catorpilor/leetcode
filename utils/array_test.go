package utils

import "testing"

func TestLowerBound(t *testing.T) {
	st := []struct {
		name   string
		arr    []int
		target int
		exp    int
	}{
		{"all identical, target is 2", []int{2, 2, 2, 2, 2, 2}, 2, 0},
		{"all identical, target not found", []int{1, 1, 1, 1, 1}, 2, 5},
		{"testcase1", []int{5, 7, 7, 8, 8, 10}, 8, 3},
		{"testcase2", []int{5, 7, 7, 8, 8, 10}, 6, 1},
		{"testcase3", []int{5, 7, 7, 8, 8, 10}, 7, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := LowerBound(tt.arr, tt.target)
			if out != tt.exp {
				t.Fatalf("with input arr:%v and target: %d wanted %d but got %d", tt.arr, tt.target, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
