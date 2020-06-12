package shuffle

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestShuffle(t *testing.T) {
	st := []struct {
		name string
		nums []int
		n    int
		exp  []int
	}{
		{"testcase1", []int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{"testcase2", []int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{"testcase3", []int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := shuf(tt.nums, tt.n)
			if diff := cmp.Diff(out, tt.exp); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
			t.Log("pass")
		})
	}
}
