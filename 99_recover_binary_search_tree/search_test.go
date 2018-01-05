package search

import (
	"reflect"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestRecoverTree(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"empty tree", nil, nil},
		{"single node", []int{1}, []int{1}},
		{"test1", []int{1, 2}, []int{2, 1}},
		{"test2", []int{3, 1, 2}, []int{2, 1, 3}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := RecoverTree(root)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
