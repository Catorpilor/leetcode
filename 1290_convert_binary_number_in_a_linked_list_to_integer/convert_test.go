package convert

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestDecimalValue(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		exp  int
	}{
		{"single node 1", utils.ConstructFromSlice([]int{1}), 1},
		{"testcase1", utils.ConstructFromSlice([]int{1, 1, 0}), 6},
		{"testcase2", utils.ConstructFromSlice([]int{0, 1, 1}), 3},
		{"testcase3", utils.ConstructFromSlice([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}), 18880},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := decimalValue(tt.head)
			if out != tt.exp {
				t.Fatalf("with input list: %s wanted %d but got %d", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestDecimalValueUseOr(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		exp  int
	}{
		{"single node 1", utils.ConstructFromSlice([]int{1}), 1},
		{"testcase1", utils.ConstructFromSlice([]int{1, 1, 0}), 6},
		{"testcase2", utils.ConstructFromSlice([]int{0, 1, 1}), 3},
		{"testcase3", utils.ConstructFromSlice([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}), 18880},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useOr(tt.head)
			if out != tt.exp {
				t.Fatalf("with input list: %s wanted %d but got %d", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
