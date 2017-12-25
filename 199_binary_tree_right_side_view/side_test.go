package side

import (
	"math"
	"reflect"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestRightSideView(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"empty", []int{}, nil},
		{"single", []int{1}, []int{1}},
		{"test1", []int{1, 2}, []int{1, 2}},
		{"test2", []int{1, 2, 3}, []int{1, 3}},
		{"test3", []int{1, 2, 3, math.MinInt32, 5, math.MinInt32, 4}, []int{1, 3, 4}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := RightSideView(root)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestRightSideView2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"empty", []int{}, nil},
		{"single", []int{1}, []int{1}},
		{"test1", []int{1, 2}, []int{1, 2}},
		{"test2", []int{1, 2, 3}, []int{1, 3}},
		{"test3", []int{1, 2, 3, math.MinInt32, 5, math.MinInt32, 4}, []int{1, 3, 4}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := RightSideView2(root)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
func TestRightSideView3(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []int
	}{
		{"empty", []int{}, nil},
		{"single", []int{1}, []int{1}},
		{"test1", []int{1, 2}, []int{1, 2}},
		{"test2", []int{1, 2, 3}, []int{1, 3}},
		{"test3", []int{1, 2, 3, math.MinInt32, 5, math.MinInt32, 4}, []int{1, 3, 4}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			root := utils.ConstructTree(c.nums)
			ret := RightSideView3(root)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
