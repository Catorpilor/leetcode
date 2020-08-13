package targetsum

import "testing"

func TestMaxSubarray(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		target int
		exp    int
	}{
		{"testcase1", []int{1, 1, 1, 1, 1}, 2, 2},
		{"testcase2", []int{-1, 3, 5, 1, 4, 2, -9}, 6, 2},
		{"testcase3", []int{-2, 6, 6, 3, 5, 4, 1, 2, 8, 10}, 10, 3},
		{"testcase4", []int{0, 0, 0}, 0, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxSubarray(tt.nums, tt.target)
			if out != tt.exp {
				t.Fatalf("with input nums:%v and target:%d wanted %d but got %d", tt.nums, tt.target, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
