package list

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestNumComps(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		g    []int
		exp  int
	}{
		{"single node", utils.ConstructFromSlice([]int{0}), []int{0}, 1},
		{"testcase1", utils.ConstructFromSlice([]int{0, 1}), []int{0}, 1},
		{"testcase2", utils.ConstructFromSlice([]int{0, 1, 2, 3, 4}), []int{0, 3, 1, 4}, 2},
		{"testcase3", utils.ConstructFromSlice([]int{0, 1, 2, 3, 4}), []int{0, 1, 2, 3, 4}, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := numComps(tt.head, tt.g)
			if out != tt.exp {
				t.Fatalf("with input list: %s and g:%v wanted %d but got %d", tt.head, tt.g, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
