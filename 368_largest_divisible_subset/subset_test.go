package subset

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLargestSubset(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"single element", []int{1}, []int{1}},
		{"all identical", []int{1, 1, 1}, []int{1, 1, 1}},
		{"testcase1", []int{1, 2, 3}, []int{2, 1}},
		{"testcase2", []int{8, 1, 4, 2}, []int{8, 4, 2, 1}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := largestSubset(tt.nums)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("with input nums:%v mismatch (-wanted,+got):\n%s", tt.nums, diff)
			}
			t.Log("pass")
		})
	}
}
