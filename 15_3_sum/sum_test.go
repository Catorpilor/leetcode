package sum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestThreeSum(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  [][]int
	}{
		{"all positive", []int{1, 2, 3}, nil},
		{"all negative", []int{-1, -2, -3}, nil},
		{"testcase1", []int{-1, 0, 1, 2, -1, -4}, [][]int{[]int{-1, 0, 1}, []int{-1, -1, 2}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := threeSum(tt.nums)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("mismatch (-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
