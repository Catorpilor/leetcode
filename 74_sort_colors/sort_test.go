package sort

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSortColors(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"single color", []int{0}, []int{0}},
		{"all identical", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{"testcase1", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.nums)
			if diff := cmp.Diff(tt.nums, tt.exp); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
			t.Log("pass")
		})
	}
}
