package mdiff

import "testing"

func TestMinDifference(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"length is less than 5", []int{1, 2, 3, 4}, 0},
		{"all equal", []int{1, 1, 1, 1, 1, 1, 1}, 0},
		{"testcase1", []int{1, 5, 0, 10, 14}, 1},
		{"testcase2", []int{6, 6, 0, 1, 1, 4, 6}, 2},
		{"testcase3", []int{1, 5, 6, 14, 15}, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minDiffs(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums:%v wanted %d but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
