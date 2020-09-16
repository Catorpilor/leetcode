package xor

import "testing"

func TestMaxResult(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"testcase1", []int{3, 10, 5, 25, 2, 8}, 28},
		{"all identical", []int{1, 1, 1, 1}, 0},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxXor(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums:%v wanted %d but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
