package discount

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFinalPrices(t *testing.T) {
	st := []struct {
		name   string
		prices []int
		exp    []int
	}{
		{"only 0ne", []int{1}, []int{1}},
		{"non descreasing", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"testcase1", []int{8, 4, 6, 2, 3}, []int{4, 2, 4, 2, 3}},
		{"testcase2", []int{10, 1, 1, 6}, []int{9, 0, 1, 6}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := finalPrices(tt.prices)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
			t.Log("pass")
		})
	}
}
