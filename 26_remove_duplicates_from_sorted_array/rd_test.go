package rd

import "testing"

func TestRemoveDups(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty nums", nil, 0},
		{"single element", []int{1}, 1},
		{"all identical", []int{1, 1, 1, 1, 1}, 1},
		{"testcase1", []int{1, 2, 3, 4, 5}, 5},
		{"testcase2", []int{1, 1, 2, 3, 3, 4}, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := removeDups(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums: %v wanted %d but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
