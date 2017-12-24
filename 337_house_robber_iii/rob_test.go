package rob

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestRob(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single", []int{2}, 2},
		{"two", []int{1, 2}, 2},
		{"test 1", []int{3, 4, 5, 1, 3, 1}, 9},
		{"test 2", []int{3, 2, 3, 4, 3, math.MinInt32, 1}, 11},
		{"failed1", []int{4, 1, math.MinInt32, 2, math.MinInt32, 3}, 7},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := Rob(root)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestRob3(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single", []int{2}, 2},
		{"two", []int{1, 2}, 2},
		{"test 1", []int{3, 4, 5, 1, 3, 1}, 9},
		{"test 2", []int{3, 2, 3, 4, 3, math.MinInt32, 1}, 11},
		{"failed1", []int{4, 1, math.MinInt32, 2, math.MinInt32, 3}, 7},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := Rob3(root)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
