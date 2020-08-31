package pancake

import (
	"testing"
)

func TestPancakeSorting(t *testing.T) {
	op := func(ans, nums []int) bool {
		return len(ans) < 10*len(nums)
	}
	st := []struct {
		name string
		arr  []int
		op   func([]int, []int) bool
	}{
		{"testcase1", []int{1, 2, 3}, op},
		{"testcase2", []int{2, 1, 3, 4}, op},
		{"testcase3", []int{3, 2, 4, 1}, op},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := pancakeSorting(tt.arr)
			if !op(out, tt.arr) {
				t.Fatalf("failed got %v", out)
			}
			t.Log("pass")
		})
	}
}
