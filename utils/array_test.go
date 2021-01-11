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

func TestUpperBound(t *testing.T) {
	st := []struct {
		name   string
		arr    []int
		target int
		exp    int
	}{
		{"empty slice", nil, 1, 0},
		{"all identical", []int{1, 1, 1}, 1, 3},
		{"testcase1", []int{1, 2, 2, 2, 3, 4}, 2, 4},
		{"testcase2", []int{2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 4}, 2, 6},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			// out := Upperbound(tt.arr, tt.target)
			out := up(tt.arr, tt.target)
			if out != tt.exp {
				t.Fatalf("with input arr:%v and target:%d wanted %d but got %d", tt.arr, tt.target, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
