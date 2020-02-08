package average

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestMaxAverage(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  float64
	}{
		{"empty slice", []int{}, 4, 0},
		{"max value", []int{1, 2, 3, 4, 5}, 1, 5.0},
		{"example", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{"k eq 5", []int{8, 0, 1, 7, 8, 6, 5, 5, 6, 7}, 5, 6.20000},
		{"k eq 6", []int{8, 0, 1, 7, 8, 6, 5, 5, 6, 7}, 6, 6.16667},
		{"k eq 7", []int{8, 0, 1, 7, 8, 6, 5, 5, 6, 7}, 7, 6.28574},
		{"k eq 8", []int{8, 0, 1, 7, 8, 6, 5, 5, 6, 7}, 8, 5.62500},
		{"k eq 9", []int{8, 0, 1, 7, 8, 6, 5, 5, 6, 7}, 9, 5.11111},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxAverage(c.nums, c.k)
			if !utils.FloatEquals(ret, c.exp) {
				t.Fatalf("expected %f but got %f with inputs %v and %d",
					c.exp, ret, c.nums, c.k)
			}
		})
	}
}
