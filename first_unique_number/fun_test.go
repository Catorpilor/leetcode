package fun

import "testing"

func initFuWithTest(nums []int, t *testing.T) *FU {
	fu := constructor(nums)
	t.Cleanup(func() {
		fu = nil
	})
	return fu
}

func TestShowFirstUnique(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"testcase1", []int{7, 7, 7, 7}, -1},
		{"testcase2", []int{2, 3, 5}, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			fu := initFuWithTest(tt.nums, t)
			out := fu.showFirstUnique()
			if out != tt.exp {
				t.Fatalf("with input nums: %v wanted %d but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
