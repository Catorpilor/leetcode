package bst

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestIsValidBST(t *testing.T) {
	st := []struct {
		name string
		nums []int // level travesal
		exp  bool
	}{
		{"empty slice", []int{}, true},
		{"single element", []int{1}, true},
		{"test12", []int{1, 2}, false},
		{"test1x2", []int{1, math.MinInt32, 2}, true},
		{"failed1", []int{10, 5, 15, math.MinInt32, math.MinInt32, 6, 20}, false},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := IsValidBST(root)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
