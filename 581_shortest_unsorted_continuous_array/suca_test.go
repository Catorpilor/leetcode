package suca

import "testing"

func TestFindUnsortedSubarray(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"single element", []int{1}, 0},
		{"all sorted", []int{1, 2, 3, 4, 5}, 0},
		{"descending order", []int{5, 4, 3, 2, 1}, 5},
		{"testcase1", []int{2, 6, 4, 8, 10, 9, 15}, 5},
		{"failed1", []int{1, 3, 4, 3, 5}, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := unsortedSubarray(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums: %v wanted %d, but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestHelper(t *testing.T) {
	st := []struct {
		name string
		nums []int
		l, r int
	}{
		{"nil slice", nil, 0, 0},
		{"single", []int{1}, 0, 0},
		{"testcase1", []int{1, 2, 3}, 0, 2},
		{"testcase2", []int{3, 2, 1}, 2, 0},
		{"testcase3", []int{2, 6, 4, 8, 10, 9, 7}, 0, 5},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			ll, rr := helper(tt.nums)
			if ll != tt.l && rr != tt.r {
				t.Fatalf("with nums: %v wanted %d & %d but got %d & %d", tt.nums, tt.l, tt.r, ll, rr)
			}
			t.Log("pass")
		})
	}
}
func TestSelectionSort(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"single element", []int{1}, 0},
		{"all sorted", []int{1, 2, 3, 4, 5}, 0},
		{"descending order", []int{5, 4, 3, 2, 1}, 5},
		{"testcase1", []int{2, 6, 4, 8, 10, 9, 15}, 5},
		{"failed1", []int{1, 3, 4, 3, 5}, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := selectionSort(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums: %v wanted %d, but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
