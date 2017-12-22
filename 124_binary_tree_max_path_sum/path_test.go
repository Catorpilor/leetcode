package path

import "testing"
import "github.com/catorpilor/leetcode/utils"

func TestMaxPathSum(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single", []int{1}, 1},
		{"level 2", []int{1, 2}, 3},
		{"test 1", []int{1, 2, 3}, 6},
		{"test 3", []int{1, 2, 3, 15, 16}, 33},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := MaxPathSum(root)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
