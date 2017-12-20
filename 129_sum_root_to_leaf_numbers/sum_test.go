package sum

import "testing"
import "github.com/catorpilor/leetcode/utils"
import "math"

func TestSumNumbers(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single", []int{1}, 1},
		{"test1", []int{1, 2, 3}, 25},
		{"test2", []int{1, 2, 3, 4}, 137},
		{"failed1", []int{4, 9, 0, math.MinInt32, 1}, 531},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := SumNumbers(root)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
