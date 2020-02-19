package plus

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestPlusOne(t *testing.T) {
	st := []struct {
		name      string
		head, exp *utils.ListNode
	}{
		{"head is nil", nil, nil},
		{"single node val < 9", utils.ConstructFromSlice([]int{3}), utils.ConstructFromSlice([]int{4})},
		{"single node val=9", utils.ConstructFromSlice([]int{9}), utils.ConstructFromSlice([]int{1, 0})},
		{"testcase1", utils.ConstructFromSlice([]int{1, 2, 3}), utils.ConstructFromSlice([]int{1, 2, 4})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := plusOne(tt.head)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list:%s wanted %s but got %s", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestPlusOneUseDummy(t *testing.T) {
	st := []struct {
		name      string
		head, exp *utils.ListNode
	}{
		{"head is nil", nil, nil},
		{"single node val < 9", utils.ConstructFromSlice([]int{3}), utils.ConstructFromSlice([]int{4})},
		{"single node val=9", utils.ConstructFromSlice([]int{9}), utils.ConstructFromSlice([]int{1, 0})},
		{"testcase1", utils.ConstructFromSlice([]int{1, 2, 3}), utils.ConstructFromSlice([]int{1, 2, 4})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useDummy(tt.head)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list:%s wanted %s but got %s", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
