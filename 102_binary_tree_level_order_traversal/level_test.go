package level

import (
	"reflect"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestLevelOrder(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  [][]int
	}{
		{"empty", []int{}, nil},
		{"single", []int{1}, [][]int{[]int{1}}},
		{"test1", []int{1, 2, 3}, [][]int{[]int{1}, []int{2, 3}}},
		{"test2", []int{1, 2, 3, 4}, [][]int{[]int{1}, []int{2, 3}, []int{4}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := LevelOrder(root)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
