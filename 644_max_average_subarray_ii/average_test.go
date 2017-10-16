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
		{"k eq 4", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{"k eq 3", []int{1, 12, -5, -6, 50, 3}, 3, 15.666667},
		{"k eq 2", []int{1, 12, -5, -6, 50, 3}, 2, 26.5},
		{"k eq 5", []int{8, 0, 1, 7, 8, 6, 5, 5, 6, 7}, 5, 6.28571},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxAverage(c.nums, c.k)
			if !utils.FloatEquals(ret, c.exp) {
				t.Fatalf("expected %f but got %f, with inputs %v and %d",
					c.exp, ret, c.nums, c.k)
			}
		})
	}
}

func TestMaxAverage2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  float64
	}{
		{"k eq 4", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{"k eq 3", []int{1, 12, -5, -6, 50, 3}, 3, 15.666667},
		{"k eq 2", []int{1, 12, -5, -6, 50, 3}, 2, 26.5},
		{"k eq 5", []int{8, 0, 1, 7, 8, 6, 5, 5, 6, 7}, 5, 6.28571},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxAverage2(c.nums, c.k)
			if !utils.FloatEquals(ret, c.exp) {
				t.Fatalf("expected %f but got %f, with inputs %v and %d",
					c.exp, ret, c.nums, c.k)
			}
		})
	}
}
