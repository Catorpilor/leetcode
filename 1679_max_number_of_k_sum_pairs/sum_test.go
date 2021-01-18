package sum

import "testing"

func TestMaxKSumOps(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  int
	}{
		{"empty nums", nil, 1, 0},
		{"single number", []int{1}, 1, 0},
		{"testcase1", []int{1, 2, 3, 4}, 5, 2},
		{"testcase2", []int{3, 1, 3, 4, 3}, 6, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxKSumOps(tt.nums, tt.k)
			if out != tt.exp {
				t.Fatalf("with input nums:%v, k:%d wanted %d but got %d", tt.nums, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
