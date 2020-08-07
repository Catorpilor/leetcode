package path

import (
	"math"
	"reflect"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestPathSum(t *testing.T) {
	st := []struct {
		name string
		nums []int
		sum  int
		exp  [][]int
	}{
		{"empty", []int{}, 2, nil},
		{"single", []int{1}, 2, nil},
		{"test1", []int{5, 4, 8}, 4, nil},
		{"test2", []int{5, 4, 8, 11, math.MinInt32, 13, 4, 7, 2, math.MinInt32, math.MinInt32, 5, 1}, 22, [][]int{
			[]int{5, 4, 11, 2}, []int{5, 8, 4, 5},
		}},
		{"failed 1", []int{-2, math.MinInt32, -3}, -5, [][]int{[]int{-2, -3}}},
		{"failed 2", []int{1, -2, -3, 1, 3, -2, math.MinInt32, -1}, 2, [][]int{[]int{1, -2, 3}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := PathSum(root, c.sum)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %v and %d",
					c.exp, ret, c.nums, c.sum)
			}
		})
	}
}

func TestPathSumUseDFS(t *testing.T) {
	st := []struct {
		name string
		nums []int
		sum  int
		exp  [][]int
	}{
		{"empty", []int{}, 2, nil},
		{"single", []int{1}, 2, nil},
		{"test1", []int{5, 4, 8}, 4, nil},
		{"test2", []int{5, 4, 8, 11, math.MinInt32, 13, 4, 7, 2, math.MinInt32, math.MinInt32, 5, 1}, 22, [][]int{
			[]int{5, 4, 11, 2}, []int{5, 8, 4, 5},
		}},
		{"failed 1", []int{-2, math.MinInt32, -3}, -5, [][]int{[]int{-2, -3}}},
		{"failed 2", []int{1, -2, -3, 1, 3, -2, math.MinInt32, -1}, 2, [][]int{[]int{1, -2, 3}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := useDFS(root, c.sum)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %v and %d",
					c.exp, ret, c.nums, c.sum)
			}
		})
	}
}
