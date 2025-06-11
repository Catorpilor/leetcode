package sub

import "testing"

func TestShortestSubArrWithAtLeastK(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  int
	}{
		{"single element", []int{1}, 1, 1},
		{"testcase1", []int{1, 2, 3}, 2, 1},
		{"testcase2", []int{2, 1, 8}, 10, 3},
		{"testcase3", []int{1, 2}, 0, 1},
		{"testcase4", []int{1, 2, 32, 21}, 55, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := shortestSubArrWithAtLeastK(tt.nums, tt.k)
			if out != tt.exp {
				t.Fatalf("with input nums:%v and k:%d wanted %d but got %d", tt.nums, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
