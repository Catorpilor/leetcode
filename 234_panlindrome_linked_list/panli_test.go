package panli

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestIsPalidrome(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		exp  bool
	}{
		{"nil list", nil, true},
		{"single element", utils.ConstructFromSlice([]int{1}), true},
		{"two elements", utils.ConstructFromSlice([]int{1, 2}), false},
		{"normal case", utils.ConstructFromSlice([]int{1, 2, 2, 1}), true},
		{"normal case 2", utils.ConstructFromSlice([]int{1, 2, 2, 3}), false},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := IsPalindrome(c.head)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %s",
					c.exp, ret, c.head)
			}
		})
	}
}
