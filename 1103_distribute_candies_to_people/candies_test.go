package candies

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDistributeCandies(t *testing.T) {
	st := []struct {
		name         string
		candies, num int
		exp          []int
	}{
		{"testcase1", 7, 4, []int{1, 2, 3, 1}},
		{"testcase2", 10, 3, []int{5, 2, 3}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := distributeCandies(tt.candies, tt.num)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
