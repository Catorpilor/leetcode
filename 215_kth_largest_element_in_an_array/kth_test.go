package kth

import "testing"

func TestKthLargest(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  int
	}{
		{"empty nums", nil, 0, 0},
		{"all identical", []int{1, 1, 1, 1, 1}, 2, 1},
		{"testcase1", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"testcase2", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := kthLargest(tt.nums, tt.k)
			if out != tt.exp {
				t.Fatalf("with input nums:%v and k:%d wanted %d but got %d", tt.nums, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestKthLargestUseSelect(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  int
	}{
		{"empty nums", nil, 0, -1},
		{"all identical", []int{1, 1, 1, 1, 1}, 2, 1},
		{"testcase1", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"testcase2", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useQuickSelect(tt.nums, tt.k)
			if out != tt.exp {
				t.Fatalf("with input nums:%v and k:%d wanted %d but got %d", tt.nums, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
