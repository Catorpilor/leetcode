package number

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSingleNumber(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"testcase1", []int{1, 2, 3, 1, 5, 2}, []int{3, 5}},
		{"testcase2", []int{1, 1, 3, 3, 2, 4, 5, 5}, []int{2, 4}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := singleNumber(tt.nums)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
